package main

import "fmt"

func calc(v int, c chan<- int) int {
	return <-c * 2
}

func main() {
	c := make(chan int, 1)
	fmt.Println("calc result is:", calc(2, c))
}
