package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/x"
)

func main() {

	rdfFile := flag.String("r", "", "Location of rdf file to load")
	flag.Parse()

	isRdf := strings.HasSuffix(*rdfFile, ".rdf")
	isRdfGz := strings.HasSuffix(*rdfFile, "rdf.gz")
	if !isRdf && !isRdfGz {
		fmt.Println("Can only use .rdf or .rdf.gz file")
		os.Exit(1)
	}
	f, err := os.Open(*rdfFile)
	x.Check(err)
	defer f.Close()
	var sc *bufio.Scanner
	if isRdfGz {
		gr, err := gzip.NewReader(f)
		x.Check(err)
		sc = bufio.NewScanner(gr)
	} else {
		sc = bufio.NewScanner(f)
	}

	//lastUID := 0
	//uids := make(map[string]uint64)

	// Load RDF
	for sc.Scan() {
		x.Check(sc.Err())

		//fmt.Printf("%s\n", sc.Text())

		nq, err := rdf.Parse(sc.Text())
		x.Check(err)
		fmt.Printf("%#v\n", nq)
	}
}
