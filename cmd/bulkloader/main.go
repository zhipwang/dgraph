package main

import (
	"bufio"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

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

var verbose = true

// TODO: Could do with some massive refactoring... E.g. run whole thing in a
// routine (or struct) and just have arg parsing in main().

func main() {

	fmt.Println()

	rdfFile := flag.String("r", "", "Location of rdf file to load")
	schemaFile := flag.String("s", "", "Location of schema file to load")
	badgerDir := flag.String("b", "", "Location of badger data directory")
	tmpDir := flag.String("tmp", os.TempDir(), "Temp directory used to use for on-disk "+
		"scratch space. Requires free space proportional to the size of the RDF file.")
	flag.Parse()

	// TODO: Handling to make sure required args have been passed.

	f, err := os.Open(*rdfFile)
	x.Check(err)
	defer f.Close()

	sc, err := rdfScanner(f, *rdfFile)
	x.Check(err)

	// TODO: Handle schema that's in gz file.
	// TODO: What is the expected file extension?

	schemaBuf, err := ioutil.ReadFile(*schemaFile)
	x.Check(err)
	schemaUpdates, err := schema.Parse(string(schemaBuf))
	x.Check(err)
	fmt.Printf("Initial schema (%d):\n", len(schemaUpdates))
	for _, sch := range schemaUpdates {
		fmt.Printf("%+v\n", sch)
	}
	fmt.Println()
	schemaStore := newSchemaStore(schemaUpdates)

	kv, err := defaultBadger(*badgerDir)
	x.Check(err)
	defer func() { x.Check(kv.Close()) }()
	// TODO: Check to make sure the badger is empty.

	plBuild := newPlBuilder(*tmpDir)
	defer plBuild.cleanUp()

	// Load RDF
	for sc.Scan() {
		x.Check(sc.Err())

		nq, err := parseNQuad(sc.Text())
		if err != nil {
			if err == rdf.ErrEmpty {
				continue
			}
			x.Check(err)
		}

		// TODO: Rename to forwardPosting and reversePosting
		p1, p2 := createEdgePostings(nq, schemaStore)

		key := x.DataKey(nq.GetPredicate(), uid(nq.GetSubject()))
		plBuild.addPosting(key, p1)

		fmt.Printf("Inserting key: %s(%d):%s\n%sValue: %+v\n\n",
			nq.GetSubject(),
			uid(nq.GetSubject()),
			nq.GetPredicate(),
			hex.Dump(key),
			p1,
		)

		if p2 != nil {
			key = x.ReverseKey(nq.GetPredicate(), uid(nq.GetObjectId()))
			plBuild.addPosting(key, p2)

			fmt.Printf("Inserting key: %s(%d):%s\n%sValue: %+v\n\n",
				nq.GetObjectId(),
				uid(nq.GetObjectId()),
				nq.GetPredicate(),
				hex.Dump(key),
				p2,
			)
		}

		key = x.DataKey("_predicate_", uid(nq.GetSubject()))
		pp := createPredicatePosting(nq.GetPredicate())
		plBuild.addPosting(key, pp)

		fmt.Printf("Inserting key: %s(%d):_predicate_\n%sValue: %+v\n\n",
			nq.GetSubject(),
			uid(nq.GetSubject()),
			hex.Dump(key),
			pp,
		)

		addIndexPostings(nq, schemaStore, plBuild)
	}

	lease(kv)

	schemaStore.write(kv)

	printUIDMap()

	plBuild.buildPostingLists(kv, schemaStore)
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

func createEdgePostings(nq gql.NQuad, ss schemaStore) (*protos.Posting, *protos.Posting) {

	// Ensure that the subject and object get UIDs.
	uid(nq.GetSubject())
	if nq.GetObjectValue() == nil {
		uid(nq.GetObjectId())
	}

	fmt.Printf("NQuad: %+v\n\n", nq.NQuad)

	de, err := nq.ToEdgeUsing(uidMap)
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

	rde, err := nq.ToEdgeUsing(uidMap)
	x.Check(err)
	rde.Entity, rde.ValueId = rde.ValueId, rde.Entity

	ss.fixEdge(rde, true)

	rp := posting.NewPosting(rde)
	rp.Op = 3

	return p, rp
}

func addIndexPostings(nq gql.NQuad, ss schemaStore, plb *plBuilder) {

	if nq.GetObjectValue() == nil {
		return // Cannot index UIDs
	}

	sch := ss.m[nq.GetPredicate()].SchemaUpdate

	for _, tokerName := range sch.GetTokenizer() {

		if verbose {
			log.Printf("[INDEX] TokenizerName: %q", tokerName)
		}

		// Find tokeniser.
		toker, ok := tok.GetTokenizer(tokerName)
		if !ok {
			x.Check(fmt.Errorf("unknown tokenizer: %v", tokerName))
		}

		// Create storage value. // TODO: Reuse the edge from create edge posting.
		de, err := nq.ToEdgeUsing(uidMap)
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
		if verbose {
			for _, tok := range toks {
				log.Printf("[INDEX] Token: %v", []byte(tok))
			}
		}

		// Store index posting.
		for _, t := range toks {
			plb.addPosting(
				x.IndexKey(nq.Predicate, t),
				&protos.Posting{
					Uid:         de.GetEntity(),
					PostingType: protos.Posting_REF,
				},
			)
		}
	}
}

func lease(kv *badger.KV) {

	// Assume that the lease is the 'next' available UID. Not yet sure how it is calculated. Seems to be blocks of 10,000.
	// E.g. 10,001, 30,001,

	// lastUID => newLease
	//    9999 => 10001
	//   10000 => 10001
	//   10001 => 10001
	//   10002 => 20001
	//   10003 => 20001
	newLease := (lastUID-2)/10000*10000 + 10001

	if verbose {
		log.Printf("[LEASE] lastUID:%d newLeaseUID:%d", lastUID, newLease)
	}

	// TODO: 10001 is hardcoded. Should be calculated dynamically.

	// TODO: Can we put this into the temp badger and do most of this function
	// automatically? Or can we somehow run 'extra' RFDs at the end? This could
	// be run just as a regular RDF.

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
	isRdf := strings.HasSuffix(filename, ".rdf")
	isRdfGz := strings.HasSuffix(filename, "rdf.gz")
	if !isRdf && !isRdfGz {
		return nil, errors.New("Can only use .rdf or .rdf.gz file")

	}
	var sc *bufio.Scanner
	if isRdfGz {
		gr, err := gzip.NewReader(f)
		x.Check(err)
		sc = bufio.NewScanner(gr)
	} else {
		sc = bufio.NewScanner(f)
	}
	return sc, nil
}

var (
	lastUID = uint64(1)
	uidMap  = map[string]uint64{}
)

func uid(str string) uint64 {
	uid, ok := uidMap[str]
	if ok {
		return uid
	}
	lastUID++
	uidMap[str] = lastUID
	return lastUID
}

func printUIDMap() {
	fmt.Println("UID map:")
	for str, uid := range uidMap {
		fmt.Printf("%d: %s\n", uid, str)
	}
	fmt.Println()
}
