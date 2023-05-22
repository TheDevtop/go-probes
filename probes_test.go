package probes

import (
	"os"
	"testing"
)

func Test_logProbe_Probe(t *testing.T) {
	p := NewLogProbe("test", os.Stderr)
	p.Probe("Test message")
}

func Test_funcProbe_Probe(t *testing.T) {
	fn := func(v ...interface{}) { println("Hello, world!") }
	p := NewFuncProbe("test", os.Stderr, fn)
	p.Probe(nil)
}

func Test_assertProbe_Probe(t *testing.T) {
	x, y := 0, 0
	p := NewAssertProbe("test", os.Stderr, NotFatal)
	p.Probe(x, y, AreEqual)

	p.Probe(x, y+1, NotEqual)
}
