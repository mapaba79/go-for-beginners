package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now() // Record start time

	urls := []string{
		"https://learnmeabitcoin.com",
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"https://this-site-does-not-exist.com",
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("[DOWN] %s - Error: %v\n", url, err)
			continue
		}
		fmt.Printf("[UP] %s - Status: %d\n", url, resp.StatusCode)
		resp.Body.Close()
	}

	fmt.Printf("Work finished in %v\n", time.Since(start))
}
