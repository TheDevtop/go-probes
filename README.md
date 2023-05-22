# (Go) Probes

These probes are here to _inform_ you, to _do stuff_ for you, and to _nuke_ your program when necessary.

```
00:01:10 [Main] The answer to life the universe and everything: 42
```

### Log Probe

The logging probe just writes the message you give it, to its file.

### Function Probe

This probe writes the fact it got probed, and procedes to execute its given function.

### Assertion Probe

This probe asserts X and Y, checks whether or not they should be equal.
And whether or not the failing of the assertion should result in a panic.
