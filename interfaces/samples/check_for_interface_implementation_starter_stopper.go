package main

type (
	Starter interface{ Start() }
	Stopper interface{ Stop() }
	Job     struct{}
)

func (Job) Start() { print("job started") }

func main() {
	var j Starter = Job{}

	if s, ok := j.(Stopper); ok {
		print("use Stopper interface\n")
		s.Stop()
	}
	print("Stopper implementation is not found.")
}
