package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	u := User{ID: 1, Name: "Bob"}
	fmt.Print(u.ID, " | ")
	fmt.Println(u.Name)
}
