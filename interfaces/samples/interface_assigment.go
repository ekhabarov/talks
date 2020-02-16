package main

// START OMIT
type Starter interface {
	Start()
}

type Job struct{}

func (Job) Start() {}

type Consumer struct{}

var s Starter
var j Job
var c Consumer

func main() {
	j = Job{}
	s = j
	s = c // error
}

// END OMIT
