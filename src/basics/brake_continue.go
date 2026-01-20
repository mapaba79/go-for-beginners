package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 2 {
			// continue: skip the rest of this turn and go to i=3
			continue
		}
		if i == 4 {
			// break: stop the whole loop right now
			break
		}
		fmt.Println("Number:", i)
	}
	// Output will only be: 1, 3
}
