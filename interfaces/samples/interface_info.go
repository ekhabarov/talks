package main

import "fmt"

// START OMIT
func main() {
	var i interface{}

	i = struct {
		Name string
	}{"Go"}

	fmt.Printf("type of i: %T\n", i)
	fmt.Printf("values with default format: %v\n", i)
	fmt.Printf("+fields: %+v\n", i)
	fmt.Printf("+types: %#v\n", i)
}

// END OMIT
