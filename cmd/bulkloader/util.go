package main

import (
	"github.com/dgraph-io/badger"
)

func defaultBadger(dir string) (*badger.KV, error) {
	opt := badger.DefaultOptions
	opt.Dir = dir
	opt.ValueDir = dir
	return badger.NewKV(&opt)
}
