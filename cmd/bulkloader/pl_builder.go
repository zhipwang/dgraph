package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

func newPlBuilder() plBuilder {
	badgerDir, err := ioutil.TempDir("", "dgraph_bulkloader")
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
			pl.Uids = bitPackUids(uids)
			plBuf, err := pl.Marshal()
			x.Check(err)
			x.Check(target.Set([]byte(k), plBuf, 0))

			// Reset for next posting list.
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
