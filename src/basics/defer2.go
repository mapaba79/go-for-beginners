package main

import "fmt"

func main() {

	// defer: this will run LAST, even though it's at the top.
	defer fmt.Println("3. Close the door")

	fmt.Println("1. Enter the room")
	fmt.Println("2. Turn on the lights")

	// When main() reaches this point, the deferred line will run.
}
