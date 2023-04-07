package main

import (
	"fmt"
)

func calc(in <-chan int, d <-chan bool, f chan<- bool) {
	sum := 0
	for {
		select {
		case <-d:
			fmt.Println("sum is:", sum)
			f <- true
			return
		case v := <-in:
			sum += v
		}
	}
}
func main() {
	c := make(chan int)
	done, finish := make(chan bool), make(chan bool)
	go calc(c, done, finish)

	for i := 0; i <= 10; i++ {
		c <- i
	}
	done <- true
	<-finish
}
