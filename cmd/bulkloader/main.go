package main

import (
	"bufio"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/gql"
	"github.com/dgraph-io/dgraph/posting"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/x"
	farm "github.com/dgryski/go-farm"
)

func main() {

	rdfFile := flag.String("r", "", "Location of rdf file to load")
	badgerDir := flag.String("b", "", "Location of badger data directory")
	tmpDir := flag.String("tmp", os.TempDir(), "Temp directory used to use for on-disk "+
		"scratch space. Requires free space proportional to the size of the RDF file.")
	flag.Parse()

	f, err := os.Open(*rdfFile)
	x.Check(err)
	defer f.Close()

	sc, err := rdfScanner(f, *rdfFile)
	x.Check(err)

	opt := badger.DefaultOptions
	opt.Dir = *badgerDir
	opt.ValueDir = *badgerDir
	kv, err := badger.NewKV(&opt)
	x.Check(err)
	defer func() { x.Check(kv.Close()) }()

	plBuild := newPlBuilder(*tmpDir)
	defer plBuild.cleanUp()

	schemaStore := newSchemaStore()

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

		schemaStore.add(nq.NQuad)

		// Ensure that the subject and object get UIDs.
		uid(nq.GetSubject())
		if nq.GetObjectValue() == nil {
			uid(nq.GetObjectId())
		}

		// TODO: Should generate key and value in their own function.
		de, err := nq.ToEdgeUsing(uidMap)
		x.Check(err)
		p := posting.NewPosting(de)
		if nq.GetObjectValue() != nil {
			// Use special sentinel UID to represent a literal node.
			p.Uid = math.MaxUint64
		}
		p.Op = 3

		key := x.DataKey(nq.GetPredicate(), uid(nq.GetSubject()))
		plBuild.addPosting(key, p)

		fmt.Printf("Inserting key: %s(%d):%s\n%sValue: %#v\n\n",
			nq.GetSubject(),
			uid(nq.GetSubject()),
			nq.GetPredicate(),
			hex.Dump(key),
			p,
		)

		key = x.DataKey("_predicate_", uid(nq.GetSubject()))
		p = createPredicatePosting(nq.GetPredicate())
		plBuild.addPosting(key, p)

		fmt.Printf("Inserting key: %s(%d):_predicate_\n%sValue: %#v\n\n",
			nq.GetSubject(),
			uid(nq.GetSubject()),
			hex.Dump(key),
			p,
		)
	}

	lease(kv)

	schemaStore.write(kv)

	printUIDMap()

	plBuild.buildPostingLists(kv)
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

func lease(kv *badger.KV) {

	// TODO: 10001 is hardcoded. Should be calculated dynamically.

	nqTmp, err := rdf.Parse("<ROOT> <_lease_> \"10001\"^^<xs:int> .")
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
		Uids:     bitPackUids([]uint64{math.MaxUint64}),
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

// TODO: Candidate for moving into pl_builder.go?
func bitPackUids(uids []uint64) []byte {
	var bp bp128.BPackEncoder
	bp.PackAppend(uids)
	buf := make([]byte, bp.Size())
	bp.WriteTo(buf)
	return buf
}
