package main

import "fmt"

func changeValue(val *int) {
	*val = 100 // Changes the value at the address
}

func main() {
	x := 10
	changeValue(&x)
	fmt.Println(x) // Prints 100
}
