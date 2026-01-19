package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	secretNumber := r.Intn(100) + 1

	fmt.Println("Guess the number between 1 and 100!")

	for {
		var guess int

		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess < secretNumber {
			fmt.Println("Too low!")
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Correct! You win!")

			break
		}
	}
}
