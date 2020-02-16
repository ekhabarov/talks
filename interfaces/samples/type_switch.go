package main

// START OMIT
type (
	Starter interface{ Start() }
	Job     struct{}
)

func (Job) Start() { print("job started\n") }

func Do(value interface{}) {
	switch v := value.(type) { // HL
	case int:
		print(5+v, "\n")
	case Starter:
		v.Start()
	case string:
		print("This is a ", v)
	}
}

func main() {
	Do(1)
	Do(Job{})
	Do("string")
}

// END OMIT
