package main

import "fmt"

func main() {
	var n, sum int

	// Ask for user input
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)

	// Calculate sum of even numbers
	for i := 2; i <= n; i += 2 {
		sum += i
	}

	// Print the sum
	fmt.Printf("Sum of even numbers from 2 to %d is %d\n", n, sum)
}
