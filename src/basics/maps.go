package main

import "fmt"

func main() {
	emails := make(map[string]string)
	emails["John"] = "john@gmail.com"
	emails["Alice"] = "alice@work.com"
	fmt.Println(emails["John"])
}
