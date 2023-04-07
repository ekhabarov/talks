package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() { // generator
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c) // indicates that there is no more values
	}()

	go func() { // consumer
		sum := 0
		for v := range c { // will work till channel will be closed.
			sum += v
		}
		fmt.Println("sum is:", sum)
		done <- true
	}()

	fmt.Println("started")
	<-done
}
