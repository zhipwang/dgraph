package main

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/x"
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

	// Load RDF
	for sc.Scan() {
		x.Check(sc.Err())

		nq, err := rdf.Parse(sc.Text())
		x.Check(err)

		fmt.Printf("%#v\n", nq)

		subject := getUid(nq.GetSubject())
		predicate := nq.GetPredicate()
		object := nq.GetObjectValue().GetDefaultVal()

		key := x.DataKey(predicate, subject)
		list := &protos.PostingList{
			Postings: []*protos.Posting{
				&protos.Posting{
					Uid:         math.MaxUint64,
					Value:       []byte(object),
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
			Uids:     bitPackUids([]uint64{math.MaxUint64}),
		}
		val, err := list.Marshal()
		x.Check(err)
		x.Check(kv.Set(key, val, 0))

		key = x.DataKey("_predicate_", subject)
		list = &protos.PostingList{
			Postings: []*protos.Posting{
				&protos.Posting{
					Uid:         407209193152762291, // TODO: Not sure where this comes from. I *think* it's the farm.Fingerprint64 of the value (i.e. predicate).
					Value:       []byte(predicate),
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
			Uids:     bitPackUids([]uint64{407209193152762291}),
		}
		val, err = list.Marshal()
		x.Check(err)
		x.Check(kv.Set(key, val, 0))
	}

	// Lease

	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], 10001)

	leaseKey := x.DataKey("_lease_", 1)
	list := &protos.PostingList{
		Postings: []*protos.Posting{
			&protos.Posting{
				Uid:         math.MaxUint64,
				Value:       buf[:],
				ValType:     protos.Posting_INT,
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
