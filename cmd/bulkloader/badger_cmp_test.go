package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

func CompareBadgers(badgerA, badgerB string) bool {

	kvA, err := defaultBadger(badgerA)
	x.Check(err)
	defer kvA.Close() // Don't check error since we're read-only.
	kvB, err := defaultBadger(badgerB)
	x.Check(err)
	defer kvB.Close() // Don't check error since we're read-only.

	itA := kvA.NewIterator(badger.DefaultIteratorOptions)
	itB := kvB.NewIterator(badger.DefaultIteratorOptions)
	itA.Seek(nil)
	itB.Seek(nil)

	cmpEq := true

	countEq := 0

	for itA.Valid() && itB.Valid() {
		itemA := itA.Item()
		itemB := itB.Item()
		keyCmp := bytes.Compare(itemA.Key(), itemB.Key())

		if keyCmp == 0 {
			if bytes.Compare(itemA.Value(), itemB.Value()) != 0 ||
				itemA.UserMeta() != itemB.UserMeta() {
				valueMismatch(itemA, itemB)
				cmpEq = false
			}
			itA.Next()
			itB.Next()
			countEq++
		} else if keyCmp < 0 {
			keyMismatch("A", itemA)
			cmpEq = false
			itA.Next()
		} else {
			keyMismatch("B", itemB)
			cmpEq = false
			itB.Next()
		}

	}
	for itA.Valid() {
		cmpEq = false
		keyMismatch("A", itA.Item())
		itA.Next()
	}
	for itB.Valid() {
		cmpEq = false
		keyMismatch("B", itB.Item())
		itB.Next()
	}

	fmt.Printf("\nEqual count: %d\n", countEq)
	return cmpEq
}

func valueMismatch(itemA, itemB *badger.KVItem) {
	fmt.Printf(`
Equal keys have different values:
K:
%vV(A) %d:
%v%vV(B) %d:
%v%v`,
		hex.Dump(itemA.Key()),
		itemA.UserMeta(),
		hex.Dump(itemA.Value()),
		niceValue(itemA.Value()),
		itemB.UserMeta(),
		hex.Dump(itemB.Value()),
		niceValue(itemB.Value()),
	)
}

func keyMismatch(label string, item *badger.KVItem) {
	fmt.Printf(`
Key present in one KV store but not the other:
K(%s):
%vV(%s) %d:
%v%v`,
		label,
		hex.Dump(item.Key()),
		label,
		item.UserMeta(),
		hex.Dump(item.Value()),
		niceValue(item.Value()),
	)
}

func niceValue(v []byte) string {

	var result string

	var pl protos.PostingList
	err := pl.Unmarshal(v)
	if err == nil {
		result += fmt.Sprintf("Pretty: %+v\n", pl)
	}

	var su protos.SchemaUpdate
	err = su.Unmarshal(v)
	if err == nil {
		result += fmt.Sprintf("Pretty: %+v\n", su)
	}

	if result == "" {
		return "Pretty: unknown conversion"
	}
	return result
}
