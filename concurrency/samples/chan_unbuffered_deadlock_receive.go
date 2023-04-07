package main

import "fmt"

func main() {
	c := make(chan int)
	<-c
	fmt.Println("received")
}
