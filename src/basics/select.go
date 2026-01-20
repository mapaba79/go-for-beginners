package main

import "fmt"

func main() {
	coffee := make(chan string)
	tea := make(chan string)

	go func() { coffee <- "Coffee is ready!" }()
	go func() { tea <- "Tea is ready!" }()

	// select: waits for the first channel that has data
	select {
	case msg1 := <-coffee:
		fmt.Println(msg1)
	case msg2 := <-tea:
		fmt.Println(msg2)
	}
}
