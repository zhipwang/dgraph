package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type progress struct {
	rdfProg     int64
	tmpKeyProg  int64
	tmpKeyTotal int64
}

func (p *progress) reportProgress() {
	start := time.Now()
	var lastRdfProg int64
	var lastPct float64
	for {
		time.Sleep(time.Second)

		rdfProg := atomic.LoadInt64(&p.rdfProg)
		tmpKeyProg := atomic.LoadInt64(&p.tmpKeyProg)
		tmpKeyTotal := atomic.LoadInt64(&p.tmpKeyTotal)

		pct := float64(tmpKeyProg) / float64(tmpKeyTotal) * 100

		elapsed := time.Since(start)
		elapsedStr := fmt.Sprintf("[%02d:%02d:%02d]",
			int(elapsed.Hours()), int(elapsed.Minutes()), int(elapsed.Seconds()))

		// TODO: Overwrite the same line each time so we don't scroll the screen.
		if tmpKeyProg == 0 {
			log.Printf("%s [Phase 1/2] RDF count: %d (%d per second)", elapsedStr, rdfProg, rdfProg-lastRdfProg)
		} else {
			log.Printf("%s [Phase 2/2] Key progress:%6.2f%% (%.3f%% per second)", elapsedStr, pct, pct-lastPct)
		}

		lastRdfProg = rdfProg
		lastPct = pct
	}

	// TODO: Summary at end
}
