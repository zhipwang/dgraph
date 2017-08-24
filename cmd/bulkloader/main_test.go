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

func TestSingleNodeWithNameAndAge(t *testing.T) {
	rdfs := `
	<peter> <name> "Peter" .

	    <peter> <age> "28"^^<xs:int> .` // Also test blank lines/weird spacing while we're here.
	runTestCaseFromString(t, rdfs)
}

func TestUpdatedValue(t *testing.T) {
	rdfs := `
	<peter> <name> "NotPeter" .
	<peter> <name> "Peter" .`
	runTestCaseFromString(t, rdfs)
}

func TestAppleIsAFruit(t *testing.T) {
	rdfs := `<apple> <is> <fruit> .`
	runTestCaseFromString(t, rdfs)
}

func TestTwoFruits(t *testing.T) {
	rdfs := `
	<apple> <is> <fruit> .
	<banana> <is> <fruit> .`
	runTestCaseFromString(t, rdfs)
}

func TestTwoFruitsWithNames(t *testing.T) {
	rdfs := `
	<apple> <is> <fruit> .
	<banana> <is> <fruit> .
	<apple> <name> "MrApple" .
	<banana> <name> "MrBanana" .`
	runTestCaseFromString(t, rdfs)
}

func TestBadSelfGeneratedSchema(t *testing.T) {
	rdfs := `
	<abc> <pred> "hello"^^<xs:string> .
	<def> <pred> "42"^^<xs:int> .`
	runTestCaseFromString(t, rdfs)
}

func TestBadSelfGeneratedSchemaReverse(t *testing.T) {
	rdfs := `
	<def> <pred> "42"^^<xs:int> .
	<abc> <pred> "hello"^^<xs:string> .`
	runTestCaseFromString(t, rdfs)
}

func TestIntConversion(t *testing.T) {
	rdfs := `
	<a> <age> "15"^^<xs:int> .
	<b> <age> "13" .`
	runTestCaseFromString(t, rdfs)
}

func TestIntConversionHex(t *testing.T) {
	rdfs := `
	<a> <age> "15"^^<xs:int> .
	<b> <age> "0xff" .`
	runTestCaseFromString(t, rdfs)
}

func TestAgeExampleFromDocos(t *testing.T) {
	rdfs := `
	<a> <age> "15"^^<xs:int> .
	<b> <age> "13" .
	<c> <age> "14"^^<xs:string> .
	<d> <age> "14.5"^^<xs:string> .
	<e> <age> "14.5" .`
	runTestCaseFromString(t, rdfs)
}

func TestSchemaMismatch(t *testing.T) {
	rdfs := `
	<s_default>  <p_default> "default" .
	<s_string>   <p_default> "str"^^<xs:string> .
	<s_dateTime> <p_default> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_date>     <p_default> "2017-08-24"^^<xs:date> .
	<s_int>      <p_default> "100"^^<xs:int> .
	<s_boolean>  <p_default> "true"^^<xs:boolean> .
	<s_double>   <p_default> "3.14159"^^<xs:double> .

	<s_string>   <p_string> "str"^^<xs:string> .
	<s_default>  <p_string> "default" .
	<s_dateTime> <p_string> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_date>     <p_string> "2017-08-24"^^<xs:date> .
	<s_int>      <p_string> "100"^^<xs:int> .
	<s_boolean>  <p_string> "true"^^<xs:boolean> .
	<s_double>   <p_string> "3.14159"^^<xs:double> .

	<s_dateTime> <p_datetime> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_datetime> "default" .
	<s_string>   <p_datetime> "str"^^<xs:string> .
	<s_date>     <p_datetime> "2017-08-24"^^<xs:date> .
	<s_int>      <p_datetime> "100"^^<xs:int> .
	<s_boolean>  <p_datetime> "true"^^<xs:boolean> .
	<s_double>   <p_datetime> "3.14159"^^<xs:double> .

	<s_date>     <p_date> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_date> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_date> "default" .
	<s_string>   <p_date> "str"^^<xs:string> .
	<s_int>      <p_date> "100"^^<xs:int> .
	<s_boolean>  <p_date> "true"^^<xs:boolean> .
	<s_double>   <p_date> "3.14159"^^<xs:double> .

	<s_int>      <p_int> "100"^^<xs:int> .
	<s_date>     <p_int> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_int> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_int> "default" .
	<s_string>   <p_int> "str"^^<xs:string> .
	<s_boolean>  <p_int> "true"^^<xs:boolean> .
	<s_double>   <p_int> "3.14159"^^<xs:double> .

	<s_boolean>  <p_boolean> "true"^^<xs:boolean> .
	<s_int>      <p_boolean> "100"^^<xs:int> .
	<s_date>     <p_boolean> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_boolean> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_boolean> "default" .
	<s_string>   <p_boolean> "str"^^<xs:string> .
	<s_double>   <p_boolean> "3.14159"^^<xs:double> .

	<s_double>   <p_double> "3.14159"^^<xs:double> .
	<s_boolean>  <p_double> "true"^^<xs:boolean> .
	<s_int>      <p_double> "100"^^<xs:int> .
	<s_date>     <p_double> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_double> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_double> "default" .
	<s_string>   <p_double> "str"^^<xs:string> .
	`
	runTestCaseFromString(t, rdfs)
}

// TODO: Indexing

// TODO: Addition of schema

// TODO: Reverse edges.

// TODO: Language.

// TODO: Some really big files.

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

// TODO: This approach is pretty nasty. Would be super ideal if the whole thing
// could just be done in process. To get things in process, would have to set
// everything up to work as a library.
//
// But then you get bad isolation when things go wrong... (e.g. x.Check kills
// the whole test not just the process). Maybe things are okay as they
// currently are.

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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		//t.Log(string(buf))
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
