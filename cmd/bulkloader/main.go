package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"os"

	"github.com/dgraph-io/dgraph/rdf"
	"github.com/dgraph-io/dgraph/x"
)

func main() {

	rdfFile := flag.String("r", "", "Location of rdf file to load")
	flag.Parse()

	//lastUID := 0
	//uids := make(map[string]uint64)

	// Load RDF
	f, err := os.Open(*rdfFile)
	x.Check(err)
	defer f.Close()
	gr, err := gzip.NewReader(f)
	x.Check(err)
	sc := bufio.NewScanner(gr)
	for sc.Scan() {
		x.Check(sc.Err())

		//fmt.Printf("%s\n", sc.Text())

		nq, err := rdf.Parse(sc.Text())
		x.Check(err)
		fmt.Printf("%#v\n", nq)
	}
}
