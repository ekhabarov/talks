package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 100
	close(c)
	fmt.Println("1st value is:", <-c)
	fmt.Println("2nd value is:", <-c)
	fmt.Println("3rd value is:", <-c)
}
