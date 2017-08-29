package main

import (
	"log"

	"github.com/dgraph-io/badger"
)

func defaultBadger(dir string) (*badger.KV, error) {
	opt := badger.DefaultOptions
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
