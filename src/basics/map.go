package main

import "fmt"

func main() {
	// map: [KeyType]ValueType
	// Here we map a country name (string) to its DDI (int)

	ddi := make(map[string]int)

	ddi["Brazil"] = 55
	ddi["Japan"] = 81

	fmt.Println("Japan's DDI is:", ddi["Japan"])
}
