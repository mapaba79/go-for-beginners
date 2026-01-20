package main

import "fmt"

func main() {
	day := "Saturday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
		// fallthrough: forces the code to run the next case
		// even if it doesn't match
		fallthrough
	case "Rest":
		fmt.Print("Time to relax.")
	default:
		// default: runs if nothing else matches
		fmt.Println("Just another day.")
	}
}
