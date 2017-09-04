package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
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

	batchSize int
}

func NewKVWriter(kv *badger.KV, prog *progress, filename string) *KVWriter {

	w := &KVWriter{
		prog:      prog,
		filename:  filename,
		batchSize: int(rand.Float64()*(10<<20) + (10 << 20)),
	}
	return w
}

func (w *KVWriter) Set(k, v []byte, meta byte) {

	w.batch = append(w.batch, entry{k, v})
	w.sz += len(k) + len(v)

	if w.sz > w.batchSize {
		w.dumpFile()
	}
}

func (w *KVWriter) Wait() {

	if w.sz > 0 {
		w.dumpFile()
	}
}

func (w *KVWriter) dumpFile() {

	w.fileCount++
	fname := fmt.Sprintf("%s_%d", w.filename, w.fileCount)
	//fmt.Printf("Writing %s\n", fname)
	//defer fmt.Printf("Finished %s\n", fname)

	atomic.AddInt64(&w.prog.outstandingSorts, 1)
	sort.Slice(w.batch, func(i, j int) bool { return bytes.Compare(w.batch[i].k, w.batch[j].k) < 0 }) // TODO Slow?
	atomic.AddInt64(&w.prog.outstandingSorts, -1)

	atomic.AddInt64(&w.prog.outstandingWrites, 1)
	fd, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	x.Check(err)

	wr := bufio.NewWriterSize(fd, 32<<20)
	for _, entry := range w.batch {
		// TODO: would also need to write sizes of keys and values to read them later.
		for i := range entry.k {
			if entry.k[i] == '\n' {
				entry.k[i] = ' '
			}
		}
		wr.Write(entry.k)
		for i := range entry.v {
			if entry.v[i] == '\n' {
				entry.v[i] = ' '
			}
		}
		wr.Write(entry.v)
		wr.WriteString("\n")
	}
	w.batch = w.batch[:0]
	x.Check(wr.Flush())
	atomic.AddInt64(&w.prog.outstandingWrites, -1)
	x.Check(fd.Close())

	w.sz = 0
}
