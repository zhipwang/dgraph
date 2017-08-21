package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/dgraph-io/badger"
)

func TestBulkLoader(t *testing.T) {
	fis, err := ioutil.ReadDir("test_data")
	noErr(t, "Could not open test_data dir:", err)
	for _, fi := range fis {
		if name := fi.Name(); len(name) >= 2 &&
			name[0] >= '0' && name[0] <= '9' &&
			name[1] >= '0' && name[1] <= '9' &&
			fi.IsDir() {
			runTestCase(t, filepath.Join("test_data", fi.Name()))
		}
	}
}

func runTestCase(t *testing.T, testDir string) {

	t.Log(testDir)

	dgraphLoaderDir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	noErr(t, "Could not create temp dir:", err)
	defer os.RemoveAll(dgraphLoaderDir)

	bulkLoaderDir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	noErr(t, "Could not create temp dir:", err)
	defer os.RemoveAll(bulkLoaderDir)

	rdfFile := filepath.Join(testDir, "data.rdf")

	loadWithDgraphLoader(t, dgraphLoaderDir, rdfFile)
	loadWithBulkLoader(t, bulkLoaderDir, rdfFile)

	cmpBadgers(t,
		filepath.Join(dgraphLoaderDir, "p"),
		filepath.Join(bulkLoaderDir, "p"),
	)
}

func loadWithDgraphLoader(t *testing.T, dataDir string, rdfFile string) {

	// The "port in use" avoidance strategy is to assign random ports.
	workerPort := randomPort()
	port := randomPort()
	grpcPort := randomPort()

	dg := exec.Command(
		"dgraph",
		"-p", filepath.Join(dataDir, "p"),
		"-w", filepath.Join(dataDir, "w"),
		"-memory_mb=1024",
		"-workerport", workerPort,
		"-port", port,
		"-grpc_port", grpcPort,
	)
	dgStdout := new(bytes.Buffer)
	dgStderr := new(bytes.Buffer)
	dg.Stdout = dgStdout
	dg.Stderr = dgStderr
	noErr(t, "Could not start dgraph:", dg.Start())

	ld := exec.Command(
		"dgraphloader",
		"-r", rdfFile,
		"-d", "localhost:"+grpcPort,
		"-cd", filepath.Join(dataDir, "c"),
	)
	ldStdout := new(bytes.Buffer)
	ldStderr := new(bytes.Buffer)
	ld.Stdout = ldStdout
	ld.Stderr = ldStderr
	noErr(t, "Could not start loader:", ld.Start())
	done := make(chan error)
	go func() { done <- ld.Wait() }()
	select {
	case err := <-done:
		if err != nil {
			t.Log(ldStdout)
			t.Log(ldStderr)
			t.Fatal("Loader error:", err)
		}
	case <-time.After(10 * time.Second):
		t.Log(ldStdout)
		t.Log(ldStderr)
		ld.Process.Kill()
		t.Fatal("Loader timed out")
	}

	noErr(t, "Couldnot signal dgraph to stop:", dg.Process.Signal(os.Interrupt))
	if err := dg.Wait(); err != nil {
		t.Log(dgStdout)
		t.Log(dgStderr)
		t.Fatal("Error after dgraph wait:", err)
	}
}

func loadWithBulkLoader(t *testing.T, dataDir string, rdfFile string) {
	noErr(t, "Could not create p dir:", os.Mkdir(filepath.Join(dataDir, "p"), 0755))

	bl := exec.Command("bulkloader",
		"-r", rdfFile,
	)
	buf, err := bl.CombinedOutput()
	if err != nil {
		t.Log(string(buf))
		t.Fatal(err)
	}
}

func cmpBadgers(t *testing.T, dgraphLoaderP, bulkLoaderP string) {

	opt := badger.DefaultOptions
	opt.Dir = dgraphLoaderP
	opt.ValueDir = opt.Dir
	wantKV, err := badger.NewKV(&opt)
	noErr(t, "Could not start want KV:", err)
	defer func() { noErr(t, "Could not close want KV:", wantKV.Close()) }()

	opt.Dir = bulkLoaderP
	opt.ValueDir = opt.Dir
	gotKV, err := badger.NewKV(&opt)
	defer func() { noErr(t, "Could not close got KV:", gotKV.Close()) }()

	wantIt := wantKV.NewIterator(badger.DefaultIteratorOptions)
	gotIt := gotKV.NewIterator(badger.DefaultIteratorOptions)
	wantIt.Seek([]byte(nil))
	gotIt.Seek([]byte(nil))

	for wantIt.Valid() && gotIt.Valid() {
		t.Fatal("ASSERT FALSE")
	}
	for wantIt.Valid() {
		t.Errorf("Exists in want but not got:\nK: %v\nV: %v\n", wantIt.Item().Key(), wantIt.Item().Value())
		wantIt.Next()
	}
	for gotIt.Valid() {
		t.Errorf("Exists in got but not want:\nK: %v\nV: %v\n", gotIt.Item().Key(), gotIt.Item().Value())
		gotIt.Next()
	}
}

func noErr(t *testing.T, msg string, err error) {
	if err != nil {
		t.Fatal(msg, err)
	}
}

func randomPort() string {
	return strconv.Itoa(rand.Intn(20000) + 20000)
}
