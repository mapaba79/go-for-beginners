package main

import "fmt"

// func: declares the function "multiply"
func multiply(a int, b int) int {
	return a * b
}

func main() {
	result := multiply(5, 19)
	fmt.Println(result)
}
