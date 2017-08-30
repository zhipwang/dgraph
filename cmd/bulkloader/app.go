package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/gql"
	"github.com/dgraph-io/dgraph/posting"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/schema"
	"github.com/dgraph-io/dgraph/tok"
	"github.com/dgraph-io/dgraph/types"
	"github.com/dgraph-io/dgraph/x"
	farm "github.com/dgryski/go-farm"
)

type options struct {
	rdfFile    string
	schemaFile string
	badgerDir  string
	tmpDir     string
}

type app struct {
	opt  options
	um   *uidMap
	ss   schemaStore
	pb   *postingListBuilder
	kv   *badger.KV
	prog *progress

	wg sync.WaitGroup // decremented to zero after all RDFs have been processed

	rdfCh   chan string
	nquadCh chan gql.NQuad
}

func newApp(opt options) (*app, error) {

	// Load schema
	schemaBuf, err := ioutil.ReadFile(opt.schemaFile)
	if err != nil {
		return nil, x.Wrapf(err, "Could not load schema.")
	}
	initialSchema, err := schema.Parse(string(schemaBuf))
	if err != nil {
		return nil, x.Wrapf(err, "Could not parse schema.")
	}

	// Create target badger.
	kv, err := defaultBadger(opt.badgerDir)
	if err != nil {
		return nil, x.Wrapf(err, "Could not create target badger.")
	}
	x.Check(err)

	prog := new(progress)
	ss := newSchemaStore(initialSchema, kv)

	a := &app{
		opt:     opt,
		um:      newUIDMap(),
		ss:      ss,
		pb:      newPostingListBuilder(opt.tmpDir, prog, kv, ss),
		kv:      kv,
		prog:    prog,
		wg:      sync.WaitGroup{},
		rdfCh:   make(chan string),    // TODO: Buffered?
		nquadCh: make(chan gql.NQuad), // TODO: Buffer?
	}

	go prog.reportProgress()
	go a.processRDFs()
	go a.processNQuads()

	return a, nil

}

func (a *app) run() {

	// TODO: Check to make sure the badger is empty.

	f, err := os.Open(a.opt.rdfFile)
	x.Checkf(err, "Could not read RDF file.")
	defer f.Close()

	var sc *bufio.Scanner
	if !strings.HasSuffix(a.opt.rdfFile, ".gz") {
		sc = bufio.NewScanner(f)
	} else {
		gzr, err := gzip.NewReader(f)
		x.Checkf(err, "Could not create gzip reader for RDF file.")
		sc = bufio.NewScanner(gzr)
	}

	a.wg.Add(1)
	for sc.Scan() {
		a.rdfCh <- sc.Text()
	}
	close(a.rdfCh)
	x.Check(sc.Err())

	a.wg.Wait()

	a.createLeaseEdge()
	a.ss.write()
	a.pb.buildPostingLists()

	a.pb.cleanUp()
	x.Check(a.kv.Close())
}

func (a *app) processRDFs() {
	for rdfLine := range a.rdfCh {
		nq, err := parseNQuad(rdfLine)
		if err != nil {
			if err == rdf.ErrEmpty {
				continue
			}
			x.Check(err)
		}
		a.nquadCh <- nq
		atomic.AddInt64(&a.prog.rdfProg, 1)
	}
	close(a.nquadCh)
}

func parseNQuad(line string) (gql.NQuad, error) {
	nq, err := rdf.Parse(line)
	if err != nil {
		return gql.NQuad{}, err
	}
	return gql.NQuad{NQuad: &nq}, nil
}

func (a *app) processNQuads() {
	for nq := range a.nquadCh {
		// TODO: Rename to forwardPosting and reversePosting
		p1, p2 := a.createEdgePostings(nq)

		countGroupHash := crc64.Checksum([]byte(nq.GetPredicate()), crc64.MakeTable(crc64.ISO))

		key := x.DataKey(nq.GetPredicate(), a.um.uid(nq.GetSubject()))
		a.pb.addPosting(key, p1, countGroupHash)

		if p2 != nil {
			key = x.ReverseKey(nq.GetPredicate(), a.um.uid(nq.GetObjectId()))
			// Reverse predicates are counted separately from normal
			// predicates, so the hash is inverted to give a separate hash.
			a.pb.addPosting(key, p2, ^countGroupHash)
		}

		key = x.DataKey("_predicate_", a.um.uid(nq.GetSubject()))
		pp := createPredicatePosting(nq.GetPredicate())
		a.pb.addPosting(key, pp, 0) // TODO: Can the _predicate_ predicate have @index(count) ?

		a.addIndexPostings(nq)
	}
	a.wg.Done()
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

func (a *app) createEdgePostings(nq gql.NQuad) (*protos.Posting, *protos.Posting) {

	// Ensure that the subject and object get UIDs.
	a.um.uid(nq.GetSubject())
	if nq.GetObjectValue() == nil {
		a.um.uid(nq.GetObjectId())
	}

	de, err := nq.ToEdgeUsing(a.um.uids)
	x.Check(err)

	a.ss.fixEdge(de, nq.ObjectValue == nil)

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
	sch := a.ss.m[nq.GetPredicate()]
	if sch.GetDirective() != protos.SchemaUpdate_REVERSE {
		return p, nil
	}

	// Reverse predicate
	x.AssertTruef(nq.GetObjectValue() == nil, "has reverse schema iff object is UID")

	rde, err := nq.ToEdgeUsing(a.um.uids)
	x.Check(err)
	rde.Entity, rde.ValueId = rde.ValueId, rde.Entity

	a.ss.fixEdge(rde, true)

	rp := posting.NewPosting(rde)
	rp.Op = 3

	return p, rp
}

func (a *app) addIndexPostings(nq gql.NQuad) {

	if nq.GetObjectValue() == nil {
		return // Cannot index UIDs
	}

	sch := a.ss.m[nq.GetPredicate()].SchemaUpdate

	for _, tokerName := range sch.GetTokenizer() {

		// Find tokeniser.
		toker, ok := tok.GetTokenizer(tokerName)
		if !ok {
			log.Fatalf("unknown tokenizer %q", tokerName)
		}

		// Create storage value. // TODO: Reuse the edge from create edge posting.
		de, err := nq.ToEdgeUsing(a.um.uids)
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
			a.pb.addPosting(
				x.IndexKey(nq.Predicate, t),
				&protos.Posting{
					Uid:         de.GetEntity(),
					PostingType: protos.Posting_REF,
				},
				0,
			)
		}
	}
}

func (a *app) createLeaseEdge() {

	newLease := a.um.lease()

	// Would be nice to be able to run this as a regular RDF, rather than as a
	// special case.

	leaseRDF := fmt.Sprintf("<ROOT> <_lease_> \"%d\"^^<xs:int> .", newLease)

	nqTmp, err := rdf.Parse(leaseRDF)
	x.Check(err)
	nq := gql.NQuad{&nqTmp}
	de, err := nq.ToEdgeUsing(map[string]uint64{"ROOT": 1})
	x.Check(err)
	p := posting.NewPosting(de)
	p.Uid = math.MaxUint64
	p.Op = 3

	leaseKey := x.DataKey(nq.GetPredicate(), de.GetEntity())
	list := &protos.PostingList{
		Postings: []*protos.Posting{p},
		Uids:     bp128.DeltaPack([]uint64{math.MaxUint64}),
	}
	val, err := list.Marshal()
	x.Check(err)
	x.Check(a.kv.Set(leaseKey, val, 0))
}

func rdfScanner(f *os.File, filename string) (*bufio.Scanner, error) {
	if !strings.HasSuffix(filename, ".gz") {
		return bufio.NewScanner(f), nil
	}
	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(gr), nil
}
