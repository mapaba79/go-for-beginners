package main

import "fmt"

func main() {
	var greeting string = "Hello"
	age := 25 // Shorthand declaration
	isGoFun := true
	fmt.Printf("%s, I am %d years old. Is go fun? %t\n",
		greeting, age, isGoFun)
}
