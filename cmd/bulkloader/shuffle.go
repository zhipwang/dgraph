package main

import (
	"bytes"

	"github.com/dgraph-io/dgraph/protos"
)

func merge(in1, in2, out chan []*protos.MapEntry) {
	ok1, me1 := <-in1
	ok2, me2 := <-in2
	for ok1 && ok2 {
		if bytes.Compare(me1.Key, me2.Key) < 0 {
			out <- me1
			me1, ok2 = <-ch1
		} else {
			out <- me2
			me2, ok2 = <-ch2
		}
	}

	if ok1 {
		out <- me1
		for me := range ch1 {
			out <- me
		}
	}
	if ok2 {
		out <- me2
		for me := range ch2 {
			out <- me
		}
	}

	close(out)
}

const internalBufferSize = 10000

func createTreeLevel(inputChs []chan *protos.MapEntry) []chan *protos.MapEntry {
	var out []chan *protos.MapEntry
	for i := 0; i+1 < len(inputChs); i += 2 {
		outCh := make(chan *protos.MapEntry, internalBufferSize)
		go merge(inputChs[i], inputChs[i+1], outCh)
		out = append(out, outCh)
	}
	if len(inputChs)%2 == 1 {
		out = append(out, inputChs[len(inputChs)-1])
	}
	return out
}

func shufflePostings(batchCh chan<- []*protos.MapEntry,
	mapEntryChs []chan *protos.MapEntry, prog *progress, ci *countIndexer) {

	// TODO
}
