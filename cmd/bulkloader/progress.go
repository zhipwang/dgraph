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

	lastRdfProg int64
	lastPct     float64

	outstandingWrites int64

	start     time.Time
	endPhase1 time.Time

	shutdown chan struct{}
}

func newProgress() *progress {
	return &progress{
		start:    time.Now(),
		shutdown: make(chan struct{}),
	}
}

func (p *progress) reportProgress() {
	for {
		select {
		case <-time.After(time.Second):
			p.report()
		case <-p.shutdown:
			p.shutdown <- struct{}{}
			break
		}
	}
}

func (p *progress) report() {
	rdfProg := atomic.LoadInt64(&p.rdfProg)
	tmpKeyProg := atomic.LoadInt64(&p.tmpKeyProg)
	tmpKeyTotal := atomic.LoadInt64(&p.tmpKeyTotal)

	pct := float64(tmpKeyProg) / float64(tmpKeyTotal) * 100

	elapsed := round(time.Since(p.start))
	elapsedStr := fmt.Sprintf("[%10s]", elapsed.String())

	// TODO: Overwrite the same line each time so we don't scroll the screen.
	if tmpKeyProg == 0 {
		fmt.Printf("%s [Phase 1/2] [RDF count: %s] [Processing speed: %s per sec] [Outstanding writes: %d]\n",
			elapsedStr,
			engNotation(float64(rdfProg)),
			engNotation(float64(rdfProg-p.lastRdfProg)),
			atomic.LoadInt64(&p.outstandingWrites),
		)
	} else {
		fmt.Printf("%s [Phase 2/2] [Key progress: %5.2f%%] [Processing Speed: %.3f%% per sec] [Outstanding writes: %d]\n",
			elapsedStr,
			pct,
			pct-p.lastPct,
			atomic.LoadInt64(&p.outstandingWrites),
		)
	}

	p.lastRdfProg = rdfProg
	p.lastPct = pct
}

func (p *progress) printSummary() {

	p.shutdown <- struct{}{}
	<-p.shutdown
	p.report()

	now := time.Now()
	phase1 := round(p.endPhase1.Sub(p.start))
	phase2 := round(now.Sub(p.endPhase1))
	total := round(now.Sub(p.start))
	fmt.Printf("Total: %v Phase1: %v Phase2: %v\n", total, phase1, phase2)
}

func round(d time.Duration) time.Duration {
	return d / 1e9 * 1e9
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
