package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "Ping!" }() // Send to channel

	msg := <-messages // Receive from channel
	fmt.Println(msg)
}
