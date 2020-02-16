package main

import "fmt"

type a struct{}
type b struct{}

func (b) String() string {
	return "String was executed"
}

func main() {
	fmt.Printf("a: %s\n", a{})
	fmt.Printf("b: %s\n", b{})
}
