// START OMIT
package main

// import ".../job"

func main() {
	var j interface{} = /*job.*/ Job{}

	if s, ok := j.(interface{ Start() }); ok {
		s.Start()
	}
}

// END OMIT

type Job struct{}

func (Job) Start() { print("job started") }
