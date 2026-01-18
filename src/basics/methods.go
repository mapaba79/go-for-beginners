package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{10, 5}
	fmt.Println("Area:", rect.Area())
}
