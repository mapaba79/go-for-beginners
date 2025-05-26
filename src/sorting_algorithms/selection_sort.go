package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []int {
	fmt.Println("Digite os números separados por espaço:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	parts := strings.Fields(text)
	numbers := make([]int, len(parts))
	for i, p := range parts {
		numbers[i], _ = strconv.Atoi(p)
	}
	return numbers
}

func printSlice(label string, arr []int) {
	fmt.Println(label, arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := readInput()
	selectionSort(arr)
	printSlice("Ordenado (Selection Sort):", arr)
}
