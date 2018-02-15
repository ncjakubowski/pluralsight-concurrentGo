package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words)) // the second param is the size of the channel buffer

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}
