package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is root")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is dog")
}

func jordan(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is jordan")
}
func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/jordan/", jordan)
	http.ListenAndServe(":8080", nil)
}

// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/"
// "/dog/"
// "/jordan/

// Add a func for each of the routes.

// Have the "/jordan/" route print out your najordan.
