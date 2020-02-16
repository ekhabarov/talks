package main

import (
	"fmt"
	"log"
	"net/http"
)

type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

func main() {
	var c Counter
	log.Fatal(http.ListenAndServe(":8044", &c))
}
