package main

import "fmt"

func main() {
	c := make(chan int, 1) // buffered

	select {
	case c <- 1:
		fmt.Println("sent")
	default:
		fmt.Println("receiver is not ready")
	}
}
