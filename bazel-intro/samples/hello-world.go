package main

import "fmt"

func main() {
	fmt.Println(`
			package main // HL

			import ( // HL
				"flag" // HL
				"fmt" // HL
			) // HL

			var name = flag.String("name", "World", "Just a name") // HL

			func main() { // HL
				flag.Parse() // HL
				fmt.Printf("Hello, %s!\n", *name) // HL
			} // HL
	`)
}
