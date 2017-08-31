package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {

	// Setting a higher number here allows more disk I/O calls to be scheduled, hence considerably
	// improving throughput. The extra CPU overhead is almost negligible in comparison. The
	// benchmark notes are located in badger-bench/randread.
	runtime.GOMAXPROCS(128)

	var opt options
	flag.BoolVar(&opt.verbose, "v", false, "Verbose")
	flag.StringVar(&opt.rdfFile, "r", "", "Location of rdf file to load")
	flag.StringVar(&opt.schemaFile, "s", "", "Location of schema file to load")
	flag.StringVar(&opt.badgerDir, "b", "", "Location of badger data directory")
	flag.StringVar(&opt.tmpDir, "tmp", os.TempDir(), "Temp directory used to use for on-disk "+
		"scratch space. Requires free space proportional to the size of the RDF file.")
	flag.IntVar(&opt.workers, "j", runtime.NumCPU()-1,
		"Number of worker threads to use (defaults to one less than logical CPUs)")
	flag.Parse()

	// TODO: Handling to make sure required args have been passed.

	app, err := newApp(opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app.run()
}
