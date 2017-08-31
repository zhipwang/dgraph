package main

import (
	"sync"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/x"
)

// TODO: The approach here isn't the simplest. A simpler approach would be to
// add to a list in `Set`, and then do the async call in `Set` itself. `Wait`
// can set any leftovers, and then wait for all asyncs to complete.

const batchSize = 1000 // TODO: Should be parameterised

type KVWriter struct {
	kv   *badger.KV
	ch   chan *badger.Entry
	wg   sync.WaitGroup
	prog *progress
}

func NewKVWriter(kv *badger.KV, prog *progress) *KVWriter {
	w := &KVWriter{
		kv:   kv,
		ch:   make(chan *badger.Entry),
		prog: prog,
	}
	go w.recvEntries()
	return w
}

func (w *KVWriter) Set(k, v []byte, meta byte) {
	w.wg.Add(1)
	w.ch <- &badger.Entry{
		Key:      k,
		Value:    v,
		UserMeta: meta,
	}
}

func (w *KVWriter) Wait() {
	close(w.ch)
	w.wg.Wait()
}

func (w *KVWriter) recvEntries() {
	es := make([]*badger.Entry, 0, batchSize)
	for e := range w.ch {
		es = append(es, e)
		if len(es) >= batchSize {
			w.setEntries(es)
			es = make([]*badger.Entry, 0, batchSize)
		}
	}
	if len(es) > 0 {
		w.setEntries(es)
	}
}

func (w *KVWriter) setEntries(entries []*badger.Entry) {
	atomic.AddInt64(&w.prog.outstandingWrites, 1)
	w.kv.BatchSetAsync(entries, func(err error) {
		x.Check(err)
		for _, e := range entries {
			x.Check(e.Error)
		}
		w.wg.Add(-len(entries))
		atomic.AddInt64(&w.prog.outstandingWrites, -1)
	})
}
