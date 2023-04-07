package main

import "fmt"

func main() {
	c := make(chan int) // unbuffered

	select {
	case c <- 1:
		fmt.Println("sent")
	default:
		fmt.Println("receiver is not ready")
	}
}
