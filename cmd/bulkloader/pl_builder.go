package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

func newPlBuilder(tmpDir string) *plBuilder {
	badgerDir, err := ioutil.TempDir(tmpDir, "dgraph_bulkloader")
	x.Check(err)
	kv, err := defaultBadger(badgerDir)
	x.Check(err)
	return &plBuilder{kv, badgerDir}
}

type plBuilder struct {
	kv        *badger.KV
	badgerDir string
}

func (b *plBuilder) cleanUp() {
	// Don't need any persistence, but still Close() anyway to close all FDs
	// before nuking the data directory.
	x.Check(b.kv.Close())
	x.Check(os.RemoveAll(b.badgerDir))
}

func (b *plBuilder) addPosting(postingListKey []byte, posting *protos.Posting) {

	var uidBuf [8]byte
	binary.BigEndian.PutUint64(uidBuf[:], posting.Uid)

	key := postingListKey
	key = append(key, uidBuf[:]...)

	var meta byte
	var val []byte
	switch posting.PostingType {
	case protos.Posting_REF:
		// val is left nil. When we read back the key/value, the UID is
		// recovered from the key.
		meta = 0x01 // Indicates posting UID rather than protos.Posting
	case protos.Posting_VALUE, protos.Posting_VALUE_LANG:
		var err error
		val, err = posting.Marshal()
		x.Check(err)
	default:
		x.AssertTruef(false, "unknown posting type")
	}

	x.Check(b.kv.Set(key, val, meta))
}

func (b *plBuilder) buildPostingLists(target *badger.KV, ss schemaStore) {

	counts := map[int][]uint64{}

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
		if iter.Item().UserMeta() == 0x01 {
			uids = append(uids, extractUID(iter.Item().Key()))
		} else {
			p := new(protos.Posting)
			err := p.Unmarshal(iter.Item().Value())
			x.Check(err)
			uids = append(uids, p.Uid)
			pl.Postings = append(pl.Postings, p)
		}

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

		parsedK := x.Parse(k)

		if verbose {
			log.Printf("[BUILD] key={%+v} finalise=%t", parsedK, finalise)
		}

		// Write posting list out to target.
		if finalise {

			if verbose {
				log.Printf("[FIN_KEY]\n%s", strings.TrimSpace(hex.Dump(k)))
			}

			// If we saw any full postings, then use a proto.PostingList as the
			// value. But include the UID-only postings in the posting list
			// (not just the proto.Posting values).

			useFullPostings := len(pl.Postings) > 0

			if useFullPostings {
				if verbose {
					log.Printf("[FIN_FULL]")
				}
				pl.Uids = bp128.DeltaPack(uids)
				plBuf, err := pl.Marshal()
				x.Check(err)
				x.Check(target.Set(k, plBuf, 0x00))
			} else {
				if verbose {
					log.Printf("[FIN_COMPACT] uids=%v", uids)
				}
				x.Check(target.Set(k, bp128.DeltaPack(uids), 0x01))
			}

			if (parsedK.IsData() || parsedK.IsReverse()) && ss.m[parsedK.Attr].GetCount() {
				cnt := len(uids)
				counts[cnt] = append(counts[cnt], parsedK.Uid)
			}

			// Reset for next posting list.
			pl.Postings = nil
			pl.Uids = nil
			uids = nil
		}

		// TODO: We're double parsing each key. With clever tracking between
		// outside of the loop, could eliminate this.
		var parsedNewK *x.ParsedKey
		if iter.Valid() {
			parsedNewK = x.Parse(newK)
		}

		if !iter.Valid() || (parsedNewK.Attr != parsedK.Attr || parsedNewK.IsReverse() != parsedK.IsReverse()) {
			// Dump out count posting lists.
			//
			// TODO: This isn't an efficient algorithm: it requires full
			// iteration over the map and max(counts) map lookups. It's
			// possible to just iterate over the map, store in a slice, and
			// fill in the gaps while iterating the slice.
			highest := -1
			for cnt := range counts {
				for i := highest + 1; i <= cnt; i++ {
					pl := counts[i]
					key := x.CountKey(parsedK.Attr, uint32(i), parsedK.IsReverse())
					if len(pl) > 0 {
						val := bp128.DeltaPack(pl)
						x.Check(target.Set(key, val, 0x01))
					} else {
						x.Check(target.Set(key, nil, 0x00))
					}
				}
				highest = cnt
			}
			counts = map[int][]uint64{} // TODO: Possibly faster to clear map while iterating. Profile to work out.
		}

		k = newK
	}
}

func extractPLKey(kvKey []byte) []byte {
	// Copy value since it's only valid until the iterator is next advanced.
	x.AssertTruef(len(kvKey) > 8, "unexpected key size")
	k := make([]byte, len(kvKey)-8)
	copy(k, kvKey)
	return k
}

func extractUID(kvKey []byte) uint64 {
	return binary.BigEndian.Uint64(kvKey[len(kvKey)-8:])
}
