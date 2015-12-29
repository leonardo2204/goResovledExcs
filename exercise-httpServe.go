package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func (s String) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w,s)
}

func main() {
	http.Handle("/string", String("Teste do ServeHTTP"))
	http.Handle("/struct", &Struct{"Ola",", ","Noob"})
	
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
