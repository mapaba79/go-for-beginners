package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You can vote.")
	} else {
		fmt.Println("Too young to vote.")
	}
}
