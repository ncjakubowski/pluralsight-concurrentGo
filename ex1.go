package main

import (
	"time"
)

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
		}
	}()

	// we put the main thread to sleep to
	// allow the inner goroutines to execute
	// and return
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
