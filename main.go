package main

import (
	"log"
	"net/http"
)


//The following code is a home handler function which writes a byte
//slice containing the message as the http response body
func home(w http.ResponseWriter, r *http.Request) {
	// Restricting the root url pattern and showing
	// a self defined 404 page
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Let's Go Snippetbox!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

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
