package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Result defines the JSON structure.
// The `json:"..."` part is a "Struct Tag".
type Result struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Time   string `json:"checked_at"`
}

func checkSite(url string, resultChan chan Result) {

	status := "Up"
	_, err := http.Get(url)

	if err != nil {
		status = "Down"
	}

	// Send a full struct through the channel instead of just a string
	resultChan <- Result{
		URL:    url,
		Status: status,
		Time:   time.Now().Format(time.RFC3339),
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.this-site-does-not-exist-123.com",
	}

	resultsChan := make(chan Result)
	var finalResults []Result // A slice to collect our results

	// Start checking
	for _, url := range urls {
		go checkSite(url, resultsChan)
	}

	// Collect results
	for i := 0; i < len(urls); i++ {
		finalResults = append(finalResults, <-resultsChan)
	}

	// --- Saving to JSON ---

	// 1. Convert the slice to JSON (Indented for human to read)
	jsonData, err := json.MarshalIndent(finalResults, "", " ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 2.  Write to a file. 0644 is the standard file permission.
	err = os.WriteFile("results.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Check complete! Results saved to results.json")
}
