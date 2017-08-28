package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

func main() {
	a := flag.String("a", "", "directory of badger A")
	b := flag.String("b", "", "directory of badger B")
	flag.Parse()
	if *a == "" || *b == "" {
		flag.Usage()
		os.Exit(1)
	}
	if !compareBadgers(*a, *b) {
		fmt.Println("Badgers not equal")
		os.Exit(1)
	}
	fmt.Println("Badgers the same!")
}

func compareBadgers(badgerA, badgerB string) bool {

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

	return cmpEq
}

func valueMismatch(itemA, itemB *badger.KVItem) {
	fmt.Printf(`
Equal keys have different values:
K:
%v%v
V(A) %d:
%v%vV(B) %d:
%v%v`,
		hex.Dump(itemA.Key()),
		niceKey(itemA.Key()),
		itemA.UserMeta(),
		hex.Dump(itemA.Value()),
		niceValue(itemA),
		itemB.UserMeta(),
		hex.Dump(itemB.Value()),
		niceValue(itemB),
	)
}

func keyMismatch(label string, item *badger.KVItem) {
	fmt.Printf(`
Key present in one KV store but not the other:
K(%s):
%v%v
V(%s) %d:
%v%v`,
		label,
		hex.Dump(item.Key()),
		niceKey(item.Key()),
		label,
		item.UserMeta(),
		hex.Dump(item.Value()),
		niceValue(item),
	)
}

func niceKey(k []byte) string {
	pk := x.Parse(k)
	return fmt.Sprintf("Pretty: %+v", pk)
}

func niceValue(item *badger.KVItem) string {

	v := item.Value()

	if item.UserMeta() == 0x01 {
		var bp bp128.BPackIterator
		bp.Init(v, 0)
		x.AssertTruef(bp.Valid(), "must be valid")
		uids := make([]uint64, bp.Length())
		bp128.DeltaUnpack(v, uids)
		return fmt.Sprintf("Pretty: uids(%d):%v\n", len(uids), uids)
	}

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

func defaultBadger(dir string) (*badger.KV, error) {
	opt := badger.DefaultOptions
	opt.Dir = dir
	opt.ValueDir = dir
	return badger.NewKV(&opt)
}
