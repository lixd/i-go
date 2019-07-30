package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// Parse flags
	sOpts, nOpts := parseFlags()
	// Force the streaming server to setup its own signal handler
	sOpts.HandleSignals = true
	// override the NoSigs for NATS since Streaming has its own signal handler
	nOpts.NoSigs = true
	// Without this option set to true, the logger is not configured.
	sOpts.EnableLogging = true
	// This will invoke RunServerWithOpts but on Windows, may run it as a service.
	if _, err := stand.Run(sOpts, nOpts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	runtime.Goexit()
}
