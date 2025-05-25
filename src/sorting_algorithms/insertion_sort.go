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

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := readInput()
	insertionSort(arr)
	printSlice("Ordenado (insertion Sort):", arr)
}
