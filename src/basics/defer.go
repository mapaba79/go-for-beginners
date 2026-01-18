package main

import "fmt"

func main() {
	defer fmt.Println("I run last!")
	fmt.Println("I run first.")

	i := 5
	defer fmt.Println("Value of i:", i)
	i++
	fmt.Println("Current i:", i)
	i *= i
	defer fmt.Println("i²:", i)
}
