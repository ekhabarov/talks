package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // buffered
	ticker := time.NewTicker(time.Second)

	select {
	case c <- 1:
		fmt.Println("sent")
	case <-ticker.C:
		fmt.Println("timeout")
	}
}
