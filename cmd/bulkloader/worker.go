package main

import (
	"hash/crc64"
	"log"
	"math"
	"sync"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/gql"
	"github.com/dgraph-io/dgraph/posting"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/tok"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/x"
	farm "github.com/dgryski/go-farm"
)

type worker struct {
	id int // for debugging

	rdfCh chan string
	wg    sync.WaitGroup

	um *uidMap
	ss *schemaStore

	tmpBadger *KVWriter

	prog *progress
}

func newWorker(
	id int,
	rdfCh chan string,
	um *uidMap,
	ss *schemaStore,
	prog *progress,
	tmpBadger *badger.KV,
) *worker {
	return &worker{
		id:    id,
		rdfCh: rdfCh,

		um: um,
		ss: ss,

		tmpBadger: NewKVWriter(tmpBadger, prog),

		prog: prog,
	}
}

func (w *worker) run() {
	w.wg.Add(1)
	for rdf := range w.rdfCh {
		w.parseRDF(rdf)
	}
	w.wg.Done()
}

func (w *worker) wait() {
	w.wg.Wait()
	w.tmpBadger.Wait()
}

func (w *worker) addPosting(key []byte, posting *protos.Posting, countGroupHash uint64) {
	w.tmpBadger.Set(packPosting(key, posting, countGroupHash))
	atomic.AddInt64(&w.prog.tmpKeyTotal, 1)
}

func (w *worker) parseRDF(rdfLine string) {
	nq, err := parseNQuad(rdfLine)
	if err != nil {
		if err == rdf.ErrEmpty {
			return
		}
		x.Checkf(err, "Could not parse RDF.")
	}

	sUID := w.um.assignUID(nq.GetSubject())
	uidM := map[string]uint64{nq.GetSubject(): sUID}
	var oUID uint64
	if nq.GetObjectValue() == nil {
		oUID = w.um.assignUID(nq.GetObjectId())
		uidM[nq.GetObjectId()] = oUID
	}

	fwdPosting, revPosting := w.createEdgePostings(nq, uidM)
	countGroupHash := crc64.Checksum([]byte(nq.GetPredicate()), crc64.MakeTable(crc64.ISO))
	key := x.DataKey(nq.GetPredicate(), sUID)
	w.addPosting(key, fwdPosting, countGroupHash)

	if revPosting != nil {
		key = x.ReverseKey(nq.GetPredicate(), oUID)
		// Reverse predicates are counted separately from normal
		// predicates, so the hash is inverted to give a separate hash.
		w.addPosting(key, revPosting, ^countGroupHash)
	}

	key = x.DataKey("_predicate_", sUID)
	pp := createPredicatePosting(nq.GetPredicate())

	w.addPosting(key, pp, 0) // TODO: Can the _predicate_ predicate have @index(count) ?

	w.addIndexPostings(nq, uidM)

	atomic.AddInt64(&w.prog.rdfProg, 1)
}

func parseNQuad(line string) (gql.NQuad, error) {
	nq, err := rdf.Parse(line)
	if err != nil {
		return gql.NQuad{}, err
	}
	return gql.NQuad{NQuad: &nq}, nil
}

func createPredicatePosting(predicate string) *protos.Posting {
	fp := farm.Fingerprint64([]byte(predicate))
	return &protos.Posting{
		Uid:         fp,
		Value:       []byte(predicate),
		ValType:     protos.Posting_DEFAULT,
		PostingType: protos.Posting_VALUE,
		Op:          3,
	}
}

func (w *worker) createEdgePostings(nq gql.NQuad, uidM map[string]uint64) (*protos.Posting, *protos.Posting) {

	de, err := nq.ToEdgeUsing(uidM)
	x.Check(err)

	w.ss.fixEdge(de, nq.ObjectValue == nil)

	p := posting.NewPosting(de)
	if nq.GetObjectValue() != nil {
		if lang := de.GetLang(); lang == "" {
			// Use special sentinel UID to represent a non-language literal node.
			p.Uid = math.MaxUint64
		} else {
			// Allow multiple versions of the same string (each with a
			// different language) by using a unique UID for each version.
			p.Uid = farm.Fingerprint64([]byte(lang))
		}
	}
	p.Op = 3

	// Early exit for no reverse edge.
	sch := w.ss.getSchema(nq.GetPredicate())
	if sch.GetDirective() != protos.SchemaUpdate_REVERSE {
		return p, nil
	}

	// Reverse predicate
	x.AssertTruef(nq.GetObjectValue() == nil, "has reverse schema iff object is UID")

	rde, err := nq.ToEdgeUsing(uidM)
	x.Check(err)
	rde.Entity, rde.ValueId = rde.ValueId, rde.Entity

	w.ss.fixEdge(rde, true)

	rp := posting.NewPosting(rde)
	rp.Op = 3

	return p, rp
}

func (w *worker) addIndexPostings(nq gql.NQuad, uidM map[string]uint64) {

	if nq.GetObjectValue() == nil {
		return // Cannot index UIDs
	}

	sch := w.ss.getSchema(nq.GetPredicate())

	for _, tokerName := range sch.GetTokenizer() {

		// Find tokeniser.
		toker, ok := tok.GetTokenizer(tokerName)
		if !ok {
			log.Fatalf("unknown tokenizer %q", tokerName)
		}

		// Create storage value. // TODO: Reuse the edge from create edge posting.
		de, err := nq.ToEdgeUsing(uidM)
		x.Check(err)
		storageVal := types.Val{
			Tid:   types.TypeID(de.GetValueType()),
			Value: de.GetValue(),
		}

		// Convert from storage type to schema type.
		var schemaVal types.Val
		schemaVal, err = types.Convert(storageVal, types.TypeID(sch.GetValueType()))
		x.Check(err) // Shouldn't error, since we've already checked for convertibility when doing edge postings.

		// Extract tokens.
		toks, err := toker.Tokens(schemaVal)

		// Store index posting.
		for _, t := range toks {
			w.tmpBadger.Set(packPosting(
				x.IndexKey(nq.Predicate, t),
				&protos.Posting{
					Uid:         de.GetEntity(),
					PostingType: protos.Posting_REF,
				},
				0,
			))
		}
	}
}
