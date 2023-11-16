package main

import (
	"log"
	"net/http"
)



func main() {

// New response multiplexer is assigned to mux and mux is said
// to be handled by the previously defined "home" function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)	
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
