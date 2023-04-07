package main

import "fmt"

func main() {
	c := make(chan int, 1) // 2nd argument is a chan size
	fmt.Println("chan value is:", <-c)
	fmt.Println("received")
}
