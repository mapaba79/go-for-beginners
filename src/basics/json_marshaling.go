package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]int{"apple": 5, "lettuce": 7}
	bytes, _ := json.Marshal(data)
	fmt.Println(string(bytes))
}
