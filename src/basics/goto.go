package main

import "fmt"

func main() {
	i := 0

NextStep: // This is a Label
	i++
	if i < 3 {
		fmt.Println("Jumping back...")
		goto NextStep // jumps back up to the label
	}
	fmt.Println("Finished!")
}
