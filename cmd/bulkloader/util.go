package main

import (
	"log"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/badger/table"
)

func defaultBadger(dir string) (*badger.KV, error) {
	opt := badger.DefaultOptions

	// Options suggested by Manish.
	opt.MapTablesTo = table.MemoryMap
	opt.ValueGCRunInterval = time.Hour * 10
	opt.SyncWrites = false

	opt.Dir = dir
	opt.ValueDir = dir
	return badger.NewKV(&opt)
}

var verbose bool = true

func LOG(format string, args ...interface{}) {
	if verbose {
		log.Printf(format, args...)
	}
}
