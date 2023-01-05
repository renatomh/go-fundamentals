package main

import (
	"log"
	"net/http"
)

// Creating a simple HTTP server

// Creating a specific function to be returned on route
func data(w http.ResponseWriter, r *http.Request) {
	// Returning the response to the request
	w.Write([]byte("Collecting data"))
}

func main() {
	// Creating routes
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// Returning the response to the request
		w.Write([]byte("Hello, web!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// Returning the response to the request
		w.Write([]byte("Loading users"))
	})

	// Using external function
	http.HandleFunc("/data", data)

	// Listening requests on specified port
	log.Fatal(http.ListenAndServe(":5000", nil))
}
