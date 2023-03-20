package probes

const (
	IsFatal  = true
	NotFatal = false
	AreEqual = true
	NotEqual = false
)

const (
	logProbeFormat       = "%s [%s] %s\n"
	funcProbeFormat      = "%s [%s] Got probed calling function.\n"
	assertProbeCorrect   = "%s [%s] Assertion is correct!\n"
	assertProbeIncorrect = "%s [%s] Assertion is not correct!\n"
	probePanicFormat     = "Probe [%s] panicked the program!"
)

const timeFormat = "15:04:04"
