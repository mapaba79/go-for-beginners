package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "localhost"

	for i := 1; i <= 65535; i++ {
		address := fmt.Sprintf("%s:%d", target, i)

		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)

		if err != nil {
			continue
		}

		conn.Close()
		fmt.Printf("Port %d is OPEN!\n", i)
	}
}
