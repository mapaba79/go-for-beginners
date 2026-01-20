package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	fmt.Println("Digite os números separados por espaço:")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n') // lê até Enter

	strNums := strings.Fields(line)
	numbers := []int{}
	for _, str := range strNums {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Erro ao converter:", str)
			return
		}
		numbers = append(numbers, num)
	}

	BubbleSort(numbers)

	fmt.Println("Números ordenados:")
	for _, v := range numbers {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
