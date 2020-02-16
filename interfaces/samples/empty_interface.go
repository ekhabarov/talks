package main

import "fmt"

// START OMIT
var (
	i interface{}
	a int
)

func main() {
	a = 10
	i = a

	fmt.Printf("a: %#v\n", a)
	fmt.Printf("i: %#v\n", i)
	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of i: %T\n", i)
}

// END OMIT
