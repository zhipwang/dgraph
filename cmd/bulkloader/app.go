package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/gql"
	"github.com/dgraph-io/dgraph/posting"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/schema"
	"github.com/dgraph-io/dgraph/x"
)

type options struct {
	verbose    bool
	rdfFile    string
	schemaFile string
	badgerDir  string
	tmpDir     string
	workers    int
}

type app struct {
	opt          options
	um           *uidMap
	ss           *schemaStore
	targetBadger *badger.KV
	prog         *progress

	rdfCh chan string

	workers []*worker

	tmpBadger    *badger.KV
	tmpBadgerDir string
}

func newApp(opt options) (*app, error) {

	// Load schema
	schemaBuf, err := ioutil.ReadFile(opt.schemaFile)
	if err != nil {
		return nil, x.Wrapf(err, "Could not load schema.")
	}
	initialSchema, err := schema.Parse(string(schemaBuf))
	if err != nil {
		return nil, x.Wrapf(err, "Could not parse schema.")
	}

	// Create target badger.
	kv, err := defaultBadger(opt.badgerDir)
	if err != nil {
		return nil, x.Wrapf(err, "Could not create target badger.")
	}
	x.Check(err)

	// Create temp badger.
	tmpBadgerDir, err := ioutil.TempDir(opt.tmpDir, "dgraph_bulkloader")
	tmpBadger, err := defaultBadger(tmpBadgerDir)
	if err != nil {
		return nil, x.Wrapf(err, "Could not create temp badger.")
	}
	x.Check(err)

	prog := newProgress()
	ss := newSchemaStore(initialSchema, kv)

	a := &app{
		opt:          opt,
		um:           newUIDMap(),
		ss:           ss,
		targetBadger: kv,
		prog:         prog,
		rdfCh:        make(chan string),
		workers:      make([]*worker, opt.workers),
		tmpBadger:    tmpBadger,
		tmpBadgerDir: tmpBadgerDir,
	}

	for i := 0; i < opt.workers; i++ {
		a.workers[i] = newWorker(i, a.rdfCh, a.um, a.ss, a.prog, a.tmpBadger)
	}
	go prog.reportProgress()

	return a, nil
}

func (a *app) run() {

	// TODO: Check to make sure the badger is empty.

	f, err := os.Open(a.opt.rdfFile)
	x.Checkf(err, "Could not read RDF file.")
	defer f.Close()

	var sc *bufio.Scanner
	if !strings.HasSuffix(a.opt.rdfFile, ".gz") {
		sc = bufio.NewScanner(f)
	} else {
		gzr, err := gzip.NewReader(f)
		x.Checkf(err, "Could not create gzip reader for RDF file.")
		sc = bufio.NewScanner(gzr)
	}

	for _, w := range a.workers {
		go w.run()
	}
	for i := 0; sc.Scan(); i++ {
		a.rdfCh <- sc.Text()
	}
	x.Check(sc.Err())
	close(a.rdfCh)
	for _, w := range a.workers {
		w.wait()
	}
	fmt.Println("[app][run] finished waiting")

	a.prog.endPhase1 = time.Now()

	a.createLeaseEdge()
	a.ss.write()
	buildPostingLists(a.tmpBadger, a.targetBadger, a.ss, a.prog)

	x.Check(a.tmpBadger.Close())
	x.Check(os.RemoveAll(a.tmpBadgerDir))
	x.Check(a.targetBadger.Close())

	fmt.Println("[app][run] before print summary")
	a.prog.printSummary()
	fmt.Println("[app][run] after print summary")

	if a.opt.verbose {
		a.um.logState()
	}

	fmt.Println("[app][run] about to return")
}

func (a *app) createLeaseEdge() {

	newLease := a.um.lease()

	// Would be nice to be able to run this as a regular RDF, rather than as a
	// special case.

	leaseRDF := fmt.Sprintf("<ROOT> <_lease_> \"%d\"^^<xs:int> .", newLease)

	nqTmp, err := rdf.Parse(leaseRDF)
	x.Check(err)
	nq := gql.NQuad{&nqTmp}
	de, err := nq.ToEdgeUsing(map[string]uint64{"ROOT": 1})
	x.Check(err)
	p := posting.NewPosting(de)
	p.Uid = math.MaxUint64
	p.Op = 3

	leaseKey := x.DataKey(nq.GetPredicate(), de.GetEntity())
	list := &protos.PostingList{
		Postings: []*protos.Posting{p},
		Uids:     bp128.DeltaPack([]uint64{math.MaxUint64}),
	}
	val, err := list.Marshal()
	x.Check(err)
	x.Check(a.targetBadger.Set(leaseKey, val, 0))
}
