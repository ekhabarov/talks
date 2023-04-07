package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello, Doe")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello, John")
}
