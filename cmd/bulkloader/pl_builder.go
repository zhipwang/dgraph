package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

func newPlBuilder(tmpDir string) plBuilder {
	badgerDir, err := ioutil.TempDir(tmpDir, "dgraph_bulkloader")
	x.Check(err)
	opt := badger.DefaultOptions
	opt.Dir = badgerDir
	opt.ValueDir = badgerDir
	kv, err := badger.NewKV(&opt)
	x.Check(err)
	return plBuilder{kv, badgerDir}
}

type plBuilder struct {
	kv        *badger.KV
	badgerDir string
}

func (b *plBuilder) cleanUp() {
	x.Check(os.RemoveAll(b.badgerDir))
}

func (b *plBuilder) addPosting(postingListKey []byte, posting *protos.Posting) {

	var uidBuf [8]byte
	binary.BigEndian.PutUint64(uidBuf[:], posting.Uid)

	key := postingListKey
	key = append(key, uidBuf[:]...)

	val, err := posting.Marshal()
	x.Check(err)

	x.Check(b.kv.Set(key, val, 0))
}

func (b *plBuilder) buildPostingLists(target *badger.KV) {

	defer func() {
		x.Check(b.kv.Close())
	}()

	pl := &protos.PostingList{}
	uids := []uint64{}
	iter := b.kv.NewIterator(badger.DefaultIteratorOptions)
	iter.Seek(nil)
	if !iter.Valid() {
		// There were no posting lists to build.
		return
	}
	k := extractPLKey(iter.Item().Key())
	for iter.Valid() {

		// Add to PL
		val := new(protos.Posting)
		err := val.Unmarshal(iter.Item().Value())
		x.Check(err)
		uids = append(uids, val.Uid)
		pl.Postings = append(pl.Postings, val)

		// Determine if we're at the end of a single posting list.
		finalise := false
		iter.Next()
		var newK []byte
		if iter.Valid() {
			newK = extractPLKey(iter.Item().Key())
			if bytes.Compare(newK, k) != 0 {
				finalise = true
			}
		} else {
			finalise = true
		}

		// Write posting list out to target.
		if finalise {

			fmt.Print("KEY:\n" + hex.Dump(k))
			fmt.Println("POSTINGS:")
			for _, p := range pl.Postings {
				fmt.Printf("%#v\n", p)
			}
			fmt.Println("END POSTINGS\n")

			x.AssertTrue(len(pl.Postings) > 0)
			// TODO: Should check to make sure all postings have the same posting type
			switch pl.Postings[0].PostingType {
			case protos.Posting_REF:
				// TODO: This is bad, since we assume the meta data is 1 (could be other things).
				x.Check(target.Set(k, bitPackUids(uids), 1))
			case protos.Posting_VALUE:
				pl.Uids = bitPackUids(uids)
				plBuf, err := pl.Marshal()
				x.Check(err)
				x.Check(target.Set(k, plBuf, 0))
			case protos.Posting_VALUE_LANG:
				x.AssertTrue(false)
			default:
				x.AssertTrue(false)
			}

			// Reset for next posting list.
			pl.Postings = nil
			pl.Uids = nil
			uids = nil
		}
		k = newK
	}
}

func extractPLKey(kvKey []byte) []byte {
	// Copy value since it's only valid until the iterator is next advanced.
	x.AssertTrue(len(kvKey) > 8)
	k := make([]byte, len(kvKey)-8)
	copy(k, kvKey)
	return k
}
