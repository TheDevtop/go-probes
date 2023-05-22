package probes

import (
	"fmt"
	"os"
	"time"
)

// The logging probe, simply writes given arguments to its file.
type logProbe struct {
	label string
	file  *os.File
}

func (p *logProbe) Probe(mesg string) {
	go fmt.Fprintf(p.file, logProbeFormat, time.Now().Format(timeFormat), p.label, mesg)
}

// The function probe, calls the given function when probed.
type funcProbe struct {
	label    string
	file     *os.File
	function func(...interface{})
}

func (p *funcProbe) Probe(objects ...interface{}) {
	go fmt.Fprintf(p.file, funcProbeFormat, time.Now().Format(timeFormat), p.label)
	go p.function(objects...)
}

// The assertion probe, asserts equality when probed.
// Warning: When fatal == true, it will panic the program.
type assertProbe struct {
	label string
	file  *os.File
	fatal bool
}

func (p *assertProbe) Probe(x, y interface{}, mustEqual bool) {
	if mustEqual && x == y {
		go fmt.Fprintf(p.file, assertProbeCorrect, time.Now().Format(timeFormat), p.label)
		return
	}
	if !mustEqual && x != y {
		go fmt.Fprintf(p.file, assertProbeCorrect, time.Now().Format(timeFormat), p.label)
		return
	}
	if fmt.Fprintf(p.file, assertProbeIncorrect, time.Now().Format(timeFormat), p.label); p.fatal {
		panic(fmt.Sprintf(probePanicFormat, p.label))
	}
}
