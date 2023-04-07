package main

import "fmt"

func main() {
	c := make(chan int, 1)
	close(c)
	if v, ok := <-c; ok {
		fmt.Println("v is:", v)
		return
	}
	fmt.Println("closed")
}
