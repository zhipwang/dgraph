package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func init() {
	for _, p := range []string{
		"github.com/dgraph-io/badger/cmd/badger_diff",
		"github.com/dgraph-io/dgraph/cmd/dgraph",
		"github.com/dgraph-io/dgraph/cmd/dgraphloader",
		"github.com/dgraph-io/dgraph/cmd/bulkloader",
	} {
		log.Printf("Installing %s", p)
		cmd := exec.Command("go", "install", p)
		buf, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(string(buf))
			log.Fatal(err)
		}
	}
}

func TestSingleNodeWithName(t *testing.T) {
	rdfs := `<peter> <name> "Peter" .`
	runTestCaseFromString(t, rdfs)
}

//func TestSingleNodeWithNameAndAge(t *testing.T) {
//rdfs := `
//<peter> <name> "Peter" .

//<peter> <age> "28"^^<xs:int> .` // Also test blank line while we're here.
//runTestCaseFromString(t, rdfs)
//}

func runTestCaseFromString(t *testing.T, rdfs string) {
	dir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	fname := filepath.Join(dir, "data.rdf")
	if err := ioutil.WriteFile(fname, []byte(rdfs), 0644); err != nil {
		t.Fatal(err)
	}
	runTestCase(t, fname)
}

func runTestCase(t *testing.T, rdfFile string) {

	dgraphLoaderDir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	noErr(t, "Could not create temp dir:", err)
	defer os.RemoveAll(dgraphLoaderDir)

	bulkLoaderDir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	noErr(t, "Could not create temp dir:", err)
	defer os.RemoveAll(bulkLoaderDir)

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
	badgerDir := filepath.Join(dataDir, "p")
	noErr(t, "Could not create p dir:", os.Mkdir(badgerDir, 0755))

	bl := exec.Command("bulkloader",
		"-r", rdfFile,
		"-b", badgerDir,
	)
	buf, err := bl.CombinedOutput()
	t.Log(string(buf))
	if err != nil {
		t.Fatal(err)
	}
}

func cmpBadgers(t *testing.T, dgraphLoaderP, bulkLoaderP string) {
	cmd := exec.Command(
		"badger_diff",
		"-a", dgraphLoaderP,
		"-b", bulkLoaderP,
	)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		t.Log(string(buf))
		t.Fatal(err)
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
