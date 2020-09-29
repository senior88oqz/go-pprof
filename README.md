# go-pprof
Go profiling example

## Runtime

The builtin profiles provided by the [runtime/pprof](https://golang.org/pkg/runtime/pprof/) package:

* **profile**: CPU profile determines where a program spends its time while actively consuming CPU cycles (as opposed while sleeping or waiting for I/O).

* **heap**: Heap profile reports the currently live allocations; used to monitor current memory usage or check for memory leaks.

* **threadcreate**: Thread creation profile reports the sections of the program that lead the creation of new OS threads.

* **goroutine**: Goroutine profile report the stack traces of all current goroutines.

* **block**: Block profile show where goroutines block waiting on synchronization primitives (including timer channels). Block profile is not enabled by default; use runtime.SetBlockProfileRate to enable it.
    > Block Profiling needs to be enabled explicitly golang.org/pkg/runtime/#SetBlockProfileRate
    > Add `runtime.SetBlockProfileRate(1)` at the top of the main function to get where goroutines block waiting on synchronization primitives

* **mutex**: Mutex profile reports the lock contentions. When you think your CPU is not fully utilized due to a mutex contention, use this profile. Mutex profile is not enabled by default, see runtime.SetMutexProfileFraction to enable.

* Profiling with `go test`
```
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

## HTTP

Open http://localhost:{listener_port}/debug/pprof/ to view all available profiles

## Refs

- https://golang.org/pkg/net/http/pprof/
- https://medium.com/compass-true-north/memory-profiling-a-go-service-cd62b90619f9
- https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/
- https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d
- https://medium.com/better-programming/profiling-a-golang-grpc-server-using-pprof-b6de1371fdd
