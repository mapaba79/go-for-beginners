package main

import "fmt"

// type: used to define a name for a new structure
// struct: a collection of different fields
type Car struct {
	Brand string
	Color string
	Year  int
}

func main() {
	myCar := Car{Brand: "Toyota", Color: "red", Year: 2024}
	fmt.Println(myCar.Brand, "-", myCar.Color, "-", myCar.Year)
}
