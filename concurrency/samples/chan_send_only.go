package main

import "fmt"

func calc(v int, c chan<- int) {
	c <- v * 2
}

func main() {
	c := make(chan int, 1)
	calc(2, c)
	fmt.Println("calc result is:", <-c)
}
