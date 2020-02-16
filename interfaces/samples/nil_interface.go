package main

import "fmt"

// START OMIT
type MyError struct{}

func (MyError) Error() string { return "MyError" }

func Do(i int) error {
	var e *MyError = nil
	if i < 0 {
		e = &MyError{}
	}
	return e
}

func main() {
	if err := Do(1); err != nil { // Is err == nil?
		fmt.Printf("err is not nil: %#v\n", err)
		return
	}
	fmt.Println("err is nil")
}

// END OMIT
