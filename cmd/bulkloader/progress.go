package main

import (
	"fmt"
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
		elapsed = elapsed / 1e9 * 1e9 // round to second
		elapsedStr := fmt.Sprintf("[%10s]", elapsed.String())

		// TODO: Overwrite the same line each time so we don't scroll the screen.
		if tmpKeyProg == 0 {
			fmt.Printf("%s [Phase 1/2] RDF count: %s  Processing speed: %s per sec\n",
				elapsedStr, engNotation(float64(rdfProg)), engNotation(float64(rdfProg-lastRdfProg)))
		} else {
			fmt.Printf("%s [Phase 2/2] Key progress: %5.2f%%  Processing Speed: %.3f%% per sec\n",
				elapsedStr, pct, pct-lastPct)
		}

		lastRdfProg = rdfProg
		lastPct = pct
	}
}

func engNotation(x float64) string {
	e := 0
	for x >= 1000 {
		x /= 1000
		e += 3
	}
	switch {
	case x >= 100:
		return fmt.Sprintf("%5.1fe%d", x, e)
	case x >= 10:
		return fmt.Sprintf("%5.2fe%d", x, e)
	default:
		return fmt.Sprintf("%5.3fe%d", x, e)
	}
}
