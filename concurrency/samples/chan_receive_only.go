package main

import "fmt"

func calc(c <-chan int) int {
	return <-c * 2
}

func main() {
	c := make(chan int, 1)
	c <- 2
	fmt.Println("calc result is:", calc(c))
}
