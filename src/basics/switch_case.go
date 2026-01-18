package main

import "fmt"

func main() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Weekend is near.")
	default:
		fmt.Println("Just another day.")
	}
}
