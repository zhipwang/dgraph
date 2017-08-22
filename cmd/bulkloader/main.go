package main

import (
	"bufio"
	"compress/gzip"
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

	// Load RDF
	for sc.Scan() {
		x.Check(sc.Err())

		nqTmp, err := rdf.Parse(sc.Text())
		x.Check(err)

		fmt.Printf("%#v\n", nqTmp)

		predicateSchema[nq.GetPredicate()] = nil

		nq := gql.NQuad{&nqTmp}
		getUid(nqTmp.GetSubject()) // TODO: Hack, this just makes sure the subject is in the map
		de, err := nq.ToEdgeUsing(uidMap)
		x.Check(err)
		p := posting.NewPosting(de)
		p.Uid = math.MaxUint64
		p.Op = 3

		key := x.DataKey(nq.GetPredicate(), getUid(nq.GetSubject()))
		list := &protos.PostingList{
			Postings: []*protos.Posting{p},
			Uids:     bitPackUids([]uint64{p.Uid}),
		}
		val, err := list.Marshal()
		x.Check(err)
		x.Check(kv.Set(key, val, 0))

		key = x.DataKey("_predicate_", getUid(nq.GetSubject()))
		fingerprint := farm.Fingerprint64([]byte(nq.GetPredicate()))
		list = &protos.PostingList{
			Postings: []*protos.Posting{
				&protos.Posting{
					Uid:         fingerprint,
					Value:       []byte(nq.GetPredicate()),
					ValType:     protos.Posting_DEFAULT,
					PostingType: protos.Posting_VALUE,
					Metadata:    nil,
					Label:       "",
					Commit:      0,
					Facets:      nil,
					Op:          3,
				},
			},
			Checksum: nil,
			Commit:   0,
			Uids:     bitPackUids([]uint64{fingerprint}),
		}
		val, err = list.Marshal()
		x.Check(err)
		x.Check(kv.Set(key, val, 0))
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
}

func lease(kv *badger.KV) {

	// TODO: 10001 is hardcoded. Should be calculated dynamically.

	nqTmp, err := rdf.Parse("<ROOT> <_lease_> \"10001\"^^<xs:int> .")
	x.Check(err)
	nq := gql.NQuad{&nqTmp}
	de, err := nq.ToEdgeUsing(map[string]uint64{"ROOT": 1}) // TODO: Should I use a global uid map?
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

func getUid(str string) uint64 {
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
