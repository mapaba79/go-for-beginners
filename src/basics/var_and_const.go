package main

import "fmt"

func main() {

	// var: can be updated later
	var speed int = 80
	speed = 100

	// const: cannot be changed after declaration
	const gravity = 9.81

	fmt.Println(speed, "|", gravity)
}
