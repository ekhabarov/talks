package main

type (
	Starter interface{ Start() }
	Stopper interface{ Stop() }
	Job     struct{}
)

func (Job) Start() { print("job started") }

func main() {
	var j interface{} = Job{}

	if s, ok := j.(Starter); ok {
		print("use Starter interface\n")
		s.Start()
	}

	if s, ok := j.(Stopper); ok {
		print("use Stopper interface\n")
		s.Stop()
	}
}
