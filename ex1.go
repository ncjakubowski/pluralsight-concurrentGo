package main

import (
	"time"
)

func main() {

	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()

	// we put the main thread to sleep to
	// allow the inner goroutines to execute
	// and return
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
