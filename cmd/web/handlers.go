package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	// Extract the value of id parameter from the query string and
	// try to convert it to an integer using the strconv.Atoi()
	// function. If it can't be converted to integer, or value is less than 1,
	// we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	//w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Using http best practices we will limit the creation
	// to methods of type POST

	// if r.Method != "POST" {
	// 	w.Header().Set("Allow","POST")
	// 	// w.WriteHeader(405)
	// 	// w.Write([]byte("Method Not Allowed"))
	// 	http.Error(w, "Method not allowed", 405)
	// 	return
	// }

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}