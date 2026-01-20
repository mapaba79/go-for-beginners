package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- url + " might be down!"
		return
	}
	c <- url + " is up and running!"
}

func main() {
	start := time.Now() // Record start time

	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"https://this-site-does-not-exist.com",
		"https://learnmeabitcoin.com",
	}

	resultsChannel := make(chan string)

	for _, url := range urls {
		// Start a Goroutine for each URL
		go checkUrl(url, resultsChannel)
	}

	// Receive the results from the channel
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-resultsChannel)
	}

	fmt.Printf("Work finished in %v\n", time.Since(start))
}
