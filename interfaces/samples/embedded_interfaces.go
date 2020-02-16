package main

type (
	Starter        interface{ Start() }
	Stopper        interface{ Stop() }
	StarterStopper interface {
		Starter
		Stopper
	}
	Job struct{}
)

func (Job) Start() { print("job started\n") }
func (Job) Stop()  { print("job stopped\n") }

func main() {
	var j interface{} = Job{}

	if s, ok := j.(StarterStopper); ok {
		print("use StarterStopper interface\n")
		s.Start()
		s.Stop()
	}
}
