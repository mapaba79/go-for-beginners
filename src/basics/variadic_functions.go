package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5))
}
