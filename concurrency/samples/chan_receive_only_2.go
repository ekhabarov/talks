package main

import "fmt"

func calc(c <-chan int) {
	v := <-c
	c <- v * 2
}

func main() {
	c := make(chan int, 1)
	c <- 2
	fmt.Println("calc result is:", <-c)
}
