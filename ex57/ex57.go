package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "%v%v %v", s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("app.devnet.imsa.edu:4000", nil))
}
