package main

import "fmt"

func main() {
	c := make(chan int, 1) // 2nd argument is a chan size
	c <- 100
	c <- 101 // deadlock happens here.
	fmt.Println("sent")
}
