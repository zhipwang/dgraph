package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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

	predicateSchema := map[string]*protos.SchemaUpdate{
		"_predicate_": nil,
		"_lease_":     &protos.SchemaUpdate{ValueType: uint32(protos.Posting_INT)},
	}

	predicates := map[string][]*protos.Posting{}

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

		fmt.Printf("%#v\n", nq.NQuad)

		predicateSchema[nq.GetPredicate()] = &protos.SchemaUpdate{
			ValueType: uint32(gql.TypeValFrom(nq.GetObjectValue()).Tid),
		}

		nq.GetObjectValue().GetVal()
		uid(nq.GetSubject()) // Ensure that the subject is in the UID map.
		de, err := nq.ToEdgeUsing(uidMap)
		x.Check(err)
		p := posting.NewPosting(de)
		p.Uid = math.MaxUint64
		p.Op = 3

		key := x.DataKey(nq.GetPredicate(), uid(nq.GetSubject()))
		list := &protos.PostingList{
			Postings: []*protos.Posting{p},
			Uids:     bitPackUids([]uint64{p.Uid}),
		}
		val, err := list.Marshal()
		x.Check(err)
		x.Check(kv.Set(key, val, 0))

		key = x.DataKey("_predicate_", uid(nq.GetSubject()))
		p = createPredicatePosting(nq.GetPredicate())
		predicates[string(key)] = append(predicates[string(key)], p)
	}

	lease(kv)

	// Schema
	for pred, sch := range predicateSchema {
		k := x.SchemaKey(pred)
		var v []byte
		if sch != nil {
			v, err = sch.Marshal()
			x.Check(err)
		}
		x.Check(kv.Set(k, v, 0))
	}

	// Postings
	for key, postings := range predicates {

		// Deduplicate
		dedup := map[uint64]*protos.Posting{}
		for _, p := range postings {
			dedup[p.Uid] = p
		}
		postings = postings[:len(dedup)]
		uids := make([]uint64, len(dedup))
		i := 0
		for uid, p := range dedup {
			postings[i] = p
			uids[i] = uid
			i++
		}

		// Create PL
		sort.Slice(postings, func(i, j int) bool { return postings[i].Uid < postings[j].Uid })
		sort.Slice(uids, func(i, j int) bool { return uids[i] < uids[j] })
		list := &protos.PostingList{
			Postings: postings,
			Uids:     bitPackUids(uids),
		}
		val, err := list.Marshal()
		x.Check(err)
		x.Check(kv.Set([]byte(key), val, 0))
	}
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

func bitPackUids(uids []uint64) []byte {
	var bp bp128.BPackEncoder
	bp.PackAppend(uids)
	buf := make([]byte, bp.Size())
	bp.WriteTo(buf)
	return buf
}
