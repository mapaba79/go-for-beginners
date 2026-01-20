package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a dummy file for testing
	fileName := "test.txt"
	content := "Go is fast.\nGo is fun.\nLearning Go now."
	os.WriteFile(fileName, []byte(content), 0644)

	// Read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)
	lines := strings.Split(text, "\n")
	words := strings.Fields(text)

	fmt.Printf("File: %s\nLines: %d\nWords: %d\nCharacters: %d\n",
		fileName, len(lines), len(words), len(text))
}
