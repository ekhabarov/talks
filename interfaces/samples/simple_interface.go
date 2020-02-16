package main

import "fmt"

// START OMIT
type Starter interface {
	Start()
}

func StartSomething(s Starter) {
	s.Start()
}

type Job struct{}

func (Job) Start() {
	fmt.Println("job started")
}

type Consumer struct{}

func (Consumer) Start() {
	fmt.Println("consumer started")
}

func main() {
	StartSomething(Job{})
	StartSomething(Consumer{})
}

// END OMIT
