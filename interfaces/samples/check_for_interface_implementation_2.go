package main

type (
	Starter interface{ Start() }
	Stopper interface{ Stop() }
	Job     struct{}
)

func (Job) Start() { print("job started") }

func main() {
	var j interface{} = Job{} // HL

	if s, ok := j.(Stopper); ok {
		s.Start() // no such method: build time error
	}
}
