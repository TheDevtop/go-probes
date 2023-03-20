package probes

import "os"

// Allocate logging probe
func NewLogProbe(label string, fd *os.File) *logProbe {
	p := new(logProbe)
	p.label = label
	p.file = fd
	return p
}

// Allocate function probe
func NewFuncProbe(label string, fd *os.File, fn func(...interface{})) *funcProbe {
	p := new(funcProbe)
	p.label = label
	p.file = fd
	p.function = fn
	return p
}

// Allocate assertion probe
func NewAssertProbe(label string, fd *os.File, fatal bool) *assertProbe {
	p := new(assertProbe)
	p.label = label
	p.file = fd
	p.fatal = fatal
	return p
}
