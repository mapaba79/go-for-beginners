package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index %d has value %d\n", index, value)
	}
}
