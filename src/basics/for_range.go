package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Mango"}

	// for + range: goes through the slice automatically
	for index, name := range fruits {
		fmt.Printf("Fruit %d is %s\n", index, name)
	}
}
