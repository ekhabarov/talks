package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("value:", <-c)
	}()

	c <- 100

	fmt.Println("sent")
}
