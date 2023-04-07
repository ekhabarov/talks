package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Concurrency")
	}()
	go func() {
		fmt.Println("In")
	}()
	go func() {
		fmt.Println("Go")
	}()

	time.Sleep(1 * time.Millisecond)
}
