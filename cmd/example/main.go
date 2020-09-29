package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/pkg/profile"
)

func main() {
	// enabled Block profiling
	runtime.SetBlockProfileRate(1)
	// optional package github.com/pkg/profile provides a simple way to manage runtime/pprof
	defer profile.Start().Stop()
	go func() {
		// open http://localhost:6060/debug/pprof/ to view all available profiles
		http.ListenAndServe("localhost:6060", nil)
	}()
}
