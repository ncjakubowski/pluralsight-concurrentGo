package main

import (
	"time"
)

func main() {
	go func() {
		println("Hello")
	}()

	go func() {
		println("Go")
	}()

	// we put the main thread to sleep to
	// allow the inner goroutines to execute
	// and return
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
