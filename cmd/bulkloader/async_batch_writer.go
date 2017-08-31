package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/x"
)

const (
	// Seem to not matter much. Tried 100, 1000, and 10000 and got same result.
	writeBatchSize = 1000
)

type KVWriter struct {
	//kv    *badger.KV
	batch []*badger.Entry
	prog  *progress
	fd    *os.File

	batchCh chan []*badger.Entry
	doneWg  sync.WaitGroup
}

func NewKVWriter(kv *badger.KV, prog *progress, filename string) *KVWriter {

	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	fmt.Printf("Writing to %s\n", filename)
	x.Check(err)

	w := &KVWriter{
		//kv:    kv,
		batch:   make([]*badger.Entry, 0, writeBatchSize),
		prog:    prog,
		fd:      fd,
		batchCh: make(chan []*badger.Entry, 2000), // can do 2000 unfinished writes at a time just like badger
	}

	go w.doWrites()

	return w
}

func (w *KVWriter) Set(k, v []byte, meta byte) {
	e := &badger.Entry{
		Key:      k,
		Value:    v,
		UserMeta: meta,
	}
	w.batch = append(w.batch, e)
	if len(w.batch) == writeBatchSize {
		w.setEntries(w.batch)
		w.batch = make([]*badger.Entry, 0, writeBatchSize)
	}
}

func (w *KVWriter) Wait() {
	w.setEntries(w.batch)
	close(w.batchCh)
	w.doneWg.Wait()
}

func (w *KVWriter) setEntries(entries []*badger.Entry) {
	atomic.AddInt64(&w.prog.outstandingWrites, 1)
	w.batchCh <- entries
}

func (w *KVWriter) doWrites() {
	w.doneWg.Add(1)
	var buf bytes.Buffer
	for entries := range w.batchCh {
		for _, e := range entries {
			buf.Write(e.Key)
			buf.Write(e.Value)
			buf.Write([]byte{e.UserMeta})
			buf.WriteString("\n")
		}
		if buf.Len() > 16<<20 {
			x.Check2(w.fd.Write(buf.Bytes()))
			buf.Reset()
		}
		atomic.AddInt64(&w.prog.outstandingWrites, int64(-len(entries)))
	}
	if buf.Len() > 0 {
		x.Check2(w.fd.Write(buf.Bytes()))
	}

	w.doneWg.Done()
}
