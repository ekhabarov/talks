package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(200 * time.Millisecond) // goroutine works a bit longer then timeout
		fmt.Println("Hello, Doe")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello, John")
}
