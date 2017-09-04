package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/x"
)

const (
// Seem to not matter much. Tried 100, 1000, and 10000 and got same result.
//writeBatchSize = 1000
)

type entry struct {
	k, v []byte
}

type KVWriter struct {
	////kv    *badger.KV
	//batch []*badger.Entry
	//prog *progress
	//fd    *os.File

	//batchCh chan []*badger.Entry
	//doneWg  sync.WaitGroup

	prog  *progress
	batch []entry
	sz    int

	fileCount int

	filename string
}

const batchSize = 100 << 20

func NewKVWriter(kv *badger.KV, prog *progress, filename string) *KVWriter {

	w := &KVWriter{
		prog:     prog,
		filename: filename,
	}
	return w
}

func (w *KVWriter) Set(k, v []byte, meta byte) {

	w.batch = append(w.batch, entry{k, v})
	w.sz += len(k) + len(v)

	if w.sz > batchSize {
		w.dumpFile()
	}
}

func (w *KVWriter) Wait() {

	if w.sz > 0 {
		w.dumpFile()
	}
}

func (w *KVWriter) dumpFile() {

	atomic.AddInt64(&w.prog.outstandingWrites, 1)
	defer atomic.AddInt64(&w.prog.outstandingWrites, -1)

	sort.Slice(w.batch, func(i, j int) bool { return bytes.Compare(w.batch[i].k, w.batch[j].k) < 0 }) // TODO Slow?

	w.fileCount++
	fd, err := os.OpenFile(fmt.Sprintf("%s_%d", w.filename, w.fileCount), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	x.Check(err)
	defer func() { x.Check(fd.Close()) }()

	wr := bufio.NewWriterSize(fd, 32<<20)
	for _, entry := range w.batch {
		// TODO: would also need to write sizes of keys and values to read them later.
		wr.Write(entry.k)
		wr.Write(entry.v)
	}
	x.Check(wr.Flush())

	w.sz = 0
}
