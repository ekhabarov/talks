package main

import "fmt"

func main() {
	go fmt.Println("Hello, Doe") // run this function without waiting a result.
	fmt.Println("Hello, John")
}
