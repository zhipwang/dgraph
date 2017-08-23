package main

import (
	"encoding/binary"
	"sort"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

type plBuilder struct {
	// TODO: Replace using a badger. Keeping in memory now for easier debugging.
	//
	// Keys: PostingList Key + UIDOfPosting
	// Values: Posting, or []byte (implies just the UID)
	inMem map[string][]byte
}

func (b *plBuilder) addPosting(postingListKey []byte, posting *protos.Posting) {
	var uidBuf [8]byte
	binary.BigEndian.PutUint64(uidBuf[:], posting.Uid)
	key := string(postingListKey) + string(uidBuf[:])
	val, err := posting.Marshal()
	x.Check(err)
	b.inMem[key] = val
}

func (b *plBuilder) buildPostingLists(target *badger.KV) {

	keys := make([]string, len(b.inMem))
	i := 0
	for kvKey := range b.inMem {
		keys[i] = kvKey
		i++
	}
	sort.Strings(keys)

	pl := &protos.PostingList{}
	uids := []uint64{}
	iter := 0
	k := extractPLKey(keys[0])
	for iter < len(keys) {

		// Add to PL
		val := new(protos.Posting)
		err := val.Unmarshal(b.inMem[keys[iter]])
		x.Check(err)
		uids = append(uids, val.Uid)
		pl.Postings = append(pl.Postings, val)

		finalise := false
		iter++
		var newK string
		if iter < len(keys) {
			newK = extractPLKey(keys[iter])
			if newK != k {
				finalise = true
			}
		} else {
			finalise = true
		}
		if finalise {
			// Write out posting list
			pl.Uids = bitPackUids(uids)
			plBuf, err := pl.Marshal()
			x.Check(err)
			x.Check(target.Set([]byte(k), plBuf, 0))

			// Reset PL for next time.
			pl.Uids = nil
			uids = nil
		}
		k = newK
	}
}
