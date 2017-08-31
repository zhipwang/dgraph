package main

import (
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/x"
)

const batchSize = 1000 // TODO: Should be parameterised

type KVWriter struct {
	kv    *badger.KV
	batch []*badger.Entry
	prog  *progress
}

func NewKVWriter(kv *badger.KV, prog *progress) *KVWriter {
	w := &KVWriter{
		kv:    kv,
		batch: make([]*badger.Entry, 0, batchSize),
		prog:  prog,
	}
	return w
}

func (w *KVWriter) Set(k, v []byte, meta byte) {
	e := &badger.Entry{
		Key:      k,
		Value:    v,
		UserMeta: meta,
	}
	w.batch = append(w.batch, e)
	if len(w.batch) == batchSize {
		w.setEntries(w.batch)
		w.batch = make([]*badger.Entry, 0, batchSize)
	}
}

func (w *KVWriter) Wait() {
	atomic.AddInt64(&w.prog.outstandingWrites, 1)
	err := w.kv.BatchSet(w.batch)
	checkErrs(err, w.batch)
	atomic.AddInt64(&w.prog.outstandingWrites, -1)
}

func (w *KVWriter) setEntries(entries []*badger.Entry) {
	atomic.AddInt64(&w.prog.outstandingWrites, 1)
	w.kv.BatchSetAsync(entries, func(err error) {
		checkErrs(err, entries)
		atomic.AddInt64(&w.prog.outstandingWrites, -1)
	})
}

func checkErrs(err error, entries []*badger.Entry) {
	x.Check(err)
	for _, e := range entries {
		x.Check(e.Error)
	}
}
