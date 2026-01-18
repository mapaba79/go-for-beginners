package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(750 * time.Millisecond)
	}
}

func main() {
	go say("Running in background") // 'go' keyword starts the routine
	say("Running in main")
}
