package main

import "fmt"

func main() {
	// chan: creates a channel that carries integers
	messages := make(chan int)

	go func() {
		messages <- 42 // Send value into channel
	}()

	value := <-messages // Receive value from channel
	fmt.Println(value)
}
