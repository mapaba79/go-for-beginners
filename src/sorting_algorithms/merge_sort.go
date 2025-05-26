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

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return append(result, append(left[i:], right[j:]...)...)
}

func main() {
	arr := readInput()
	sorted := mergeSort(arr)
	printSlice("Ordenado (Merge Sort):", sorted)
}
