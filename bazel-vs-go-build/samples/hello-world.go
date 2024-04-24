package main

import "fmt"

func main() {
	fmt.Println(`
			package main // HL

			import "fmt" // HL

			func main() { // HL
				fmt.Println("Hello, World!") // HL
			} // HL
	`)
}
