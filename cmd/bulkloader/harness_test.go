package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func init() {
	// Need some binaries installed to make these tests work.  Easier to
	// install these programmatically to avoid problems with accidentally out
	// of date binaries.

	// TODO: When running these binaries, we currently assume they are in PATH.
	// This is bad because they might not be in PATH. Or even worse, something
	// else with the same name could be in PATH shaddowing the expected
	// version.

	for _, p := range []string{
		"github.com/dgraph-io/dgraph/cmd/dgraph",
		"github.com/dgraph-io/dgraph/cmd/dgraphloader",
		"github.com/dgraph-io/dgraph/cmd/bulkloader",
	} {
		fmt.Printf("Installing %s\n", p)
		cmd := exec.Command("go", "install", p)
		buf, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(buf))
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func runTestCaseFromString(t *testing.T, schema, rdfs string) {
	dir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	rdfFilename := filepath.Join(dir, "data.rdf")
	if err := ioutil.WriteFile(rdfFilename, []byte(rdfs), 0644); err != nil {
		t.Fatal(err)
	}
	schemaFilename := filepath.Join(dir, "schema.json")
	if err := ioutil.WriteFile(schemaFilename, []byte(schema), 0644); err != nil {
		t.Fatal(err)
	}

	runTestCase(t, rdfFilename, schemaFilename)
}

// TODO: This approach is pretty nasty. Would be super ideal if the whole thing
// could just be done in process. To get things in process, would have to set
// everything up to work as a library.
//
// But then you get bad isolation when things go wrong... (e.g. x.Check kills
// the whole test not just the process). Maybe things are okay as they
// currently are.

func runTestCase(t *testing.T, rdfFile string, schemaFile string) {

	dgraphLoaderDir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	noErr(t, "Could not create temp dir:", err)
	defer os.RemoveAll(dgraphLoaderDir)

	bulkLoaderDir, err := ioutil.TempDir("", "dgraph_bulk_loader_test")
	noErr(t, "Could not create temp dir:", err)
	defer os.RemoveAll(bulkLoaderDir)

	loadWithDgraphLoader(t, dgraphLoaderDir, rdfFile, schemaFile)
	loadWithBulkLoader(t, bulkLoaderDir, rdfFile, schemaFile)

	if !CompareBadgers(
		filepath.Join(dgraphLoaderDir, "p"),
		filepath.Join(bulkLoaderDir, "p"),
	) {
		t.Fatal("Badgers didn't compare equal")
	}
}

func loadWithDgraphLoader(t *testing.T, dataDir string, rdfFile, schemaFile string) {

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

	// Wait a short amount of time for dgraph to start listening for gRPC.
	time.Sleep(1000 * time.Millisecond)

	ld := exec.Command(
		"dgraphloader",
		"-r", rdfFile,
		"-d", "localhost:"+grpcPort,
		"-cd", filepath.Join(dataDir, "c"),
		"-s", schemaFile,
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

	noErr(t, "Couldn't signal dgraph to stop:", dg.Process.Signal(os.Interrupt))
	if err := dg.Wait(); err != nil {
		t.Log(dgStdout)
		t.Log(dgStderr)
		t.Fatal("Error after dgraph wait:", err)
	}
}

func loadWithBulkLoader(t *testing.T, dataDir string, rdfFile string, schemaFile string) {
	badgerDir := filepath.Join(dataDir, "p")
	noErr(t, "Could not create p dir:", os.Mkdir(badgerDir, 0755))

	bl := exec.Command("bulkloader",
		"-r", rdfFile,
		"-b", badgerDir,
		"-s", schemaFile,
	)
	buf, err := bl.CombinedOutput()
	t.Log(string(buf))
	if err != nil {
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
