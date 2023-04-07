package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		fmt.Println("value:", <-c) // 1. blocked, till sender will send something.
		done <- true               // 3. blocked, till receiver will be ready.
	}()

	c <- 100 // 1. blocked, till receiver will be ready.

	fmt.Println("sent") // 2. print
	<-done              // 3. blocked, till sender will be ready.
}
