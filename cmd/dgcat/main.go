package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/x"
)

func main() {
	a := flag.String("a", "", "directory of badger")
	flag.Parse()
	if *a == "" {
		flag.Usage()
		os.Exit(1)
	}

	badger.DefaultOptions.Dir = *a
	badger.DefaultOptions.ValueDir = *a
	kv, err := badger.NewKV(&badger.DefaultOptions)
	x.Check(err)

	iter := kv.NewIterator(badger.DefaultIteratorOptions)
	iter.Rewind()
	iter.Seek(nil)

	for iter.Valid() {
		//fmt.Println("K")
		fmt.Printf("%s\n", hex.Dump(iter.Item().Key()))
		iter.Next()
	}
}
