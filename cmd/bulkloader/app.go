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

func run(opt options) error {
	f, err := os.Open(opt.rdfFile)
	if err != nil {
		log.Fatalf("could not read rdf file: %v", err)
	}
	defer f.Close()

	sc, err := rdfScanner(f, opt.rdfFile)
	x.Check(err)

	// TODO: Handle schema that's in gz file.
	// TODO: What is the expected file extension?

	schemaBuf, err := ioutil.ReadFile(opt.schemaFile)
	if err != nil {
		log.Fatalf("could not read schema file: %v", err)
	}
	schemaUpdates, err := schema.Parse(string(schemaBuf))
	x.Check(err)
	schemaStore := newSchemaStore(schemaUpdates)

	var um = newUIDMap()

	kv, err := defaultBadger(opt.badgerDir)
	x.Check(err)
	defer func() { x.Check(kv.Close()) }()
	// TODO: Check to make sure the badger is empty.

	prog := progress{}
	go prog.reportProgress()

	plBuild := newPlBuilder(opt.tmpDir, &prog)
	defer plBuild.cleanUp()

	// Load RDF
	for sc.Scan() {
		x.Check(sc.Err())

		atomic.AddInt64(&prog.rdfProg, 1)

		nq, err := parseNQuad(sc.Text())
		if err != nil {
			if err == rdf.ErrEmpty {
				continue
			}
			x.Check(err)
		}

		// TODO: Rename to forwardPosting and reversePosting
		p1, p2 := createEdgePostings(nq, schemaStore, um)

		countGroupHash := crc64.Checksum([]byte(nq.GetPredicate()), crc64.MakeTable(crc64.ISO))

		key := x.DataKey(nq.GetPredicate(), um.uid(nq.GetSubject()))
		plBuild.addPosting(key, p1, countGroupHash)

		if p2 != nil {
			key = x.ReverseKey(nq.GetPredicate(), um.uid(nq.GetObjectId()))
			// Reverse predicates are counted separately from normal
			// predicates, so the hash is inverted to give a separate hash.
			plBuild.addPosting(key, p2, ^countGroupHash)
		}

		key = x.DataKey("_predicate_", um.uid(nq.GetSubject()))
		pp := createPredicatePosting(nq.GetPredicate())
		plBuild.addPosting(key, pp, 0) // TODO: Can the _predicate_ predicate have @index(count) ?

		addIndexPostings(nq, schemaStore, plBuild, um)
	}

	lease(kv, um)

	schemaStore.write(kv)

	plBuild.buildPostingLists(kv, schemaStore)

	return nil
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

func createEdgePostings(nq gql.NQuad, ss schemaStore, um *uidMap) (*protos.Posting, *protos.Posting) {

	// Ensure that the subject and object get UIDs.
	um.uid(nq.GetSubject())
	if nq.GetObjectValue() == nil {
		um.uid(nq.GetObjectId())
	}

	de, err := nq.ToEdgeUsing(um.uids)
	x.Check(err)

	ss.fixEdge(de, nq.ObjectValue == nil)

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
	sch := ss.m[nq.GetPredicate()]
	if sch.GetDirective() != protos.SchemaUpdate_REVERSE {
		return p, nil
	}

	// Reverse predicate
	x.AssertTruef(nq.GetObjectValue() == nil, "has reverse schema iff object is UID")

	rde, err := nq.ToEdgeUsing(um.uids)
	x.Check(err)
	rde.Entity, rde.ValueId = rde.ValueId, rde.Entity

	ss.fixEdge(rde, true)

	rp := posting.NewPosting(rde)
	rp.Op = 3

	return p, rp
}

func addIndexPostings(nq gql.NQuad, ss schemaStore, plb *plBuilder, um *uidMap) {

	if nq.GetObjectValue() == nil {
		return // Cannot index UIDs
	}

	sch := ss.m[nq.GetPredicate()].SchemaUpdate

	for _, tokerName := range sch.GetTokenizer() {

		// Find tokeniser.
		toker, ok := tok.GetTokenizer(tokerName)
		if !ok {
			log.Fatalf("unknown tokenizer %q", tokerName)
		}

		// Create storage value. // TODO: Reuse the edge from create edge posting.
		de, err := nq.ToEdgeUsing(um.uids)
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
			plb.addPosting(
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

func lease(kv *badger.KV, um *uidMap) {

	newLease := um.lease()

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
	x.Check(kv.Set(leaseKey, val, 0))
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
