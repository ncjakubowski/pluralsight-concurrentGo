package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1) //channel that accepts strings
	ch <- "Hello"
	fmt.Println(<-ch) // drain the channel using the receive operator
}
