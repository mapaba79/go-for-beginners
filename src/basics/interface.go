package main

import "fmt"

// interface: defines a "contract"
type Speaker interface {
	Speak() string
}

type Dog struct{}

// This method makes Dog "satisfy" the Speaker interface
func (d Dog) Speak() string {
	return "Woof"
}

func main() {
	var s Speaker = Dog{}
	fmt.Println(s.Speak())
}
