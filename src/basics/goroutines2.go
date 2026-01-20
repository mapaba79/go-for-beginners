package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from the background!")
}

func main() {

	// go: runs this function without waiting for it to finish
	go sayHello()

	fmt.Println("Hello from the main function!")

	// Wait a bit so the background worker can finish
	time.Sleep(2 * time.Second)
}
