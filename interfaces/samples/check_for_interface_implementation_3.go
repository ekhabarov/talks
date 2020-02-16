package main

type (
	Starter interface{ Start() }
	Job     struct{}
)

func (Job) Start() { print("job started") }

func main() {
	var j Job = Job{} // HL

	if s, ok := j.(Starter); ok {
		print("use Starter interface\n")
		s.Start()
	}
}
