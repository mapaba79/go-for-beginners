package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5) // Adding elements
	fmt.Println(numbers)
}
