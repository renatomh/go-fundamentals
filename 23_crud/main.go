package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Sample on how to CREATE, READ, UPDATE and DELETE data with Go
		* CREATE 	- 	POST
		* READ 		- 	GET
		* UPDATE 	- 	PUT/PATCH
		* DELETE 	- 	DELETE
*/

func main() {
	// Creating the router
	router := mux.NewRouter()

	// Function to handle user routes
	// We can specify the accepted methods
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	// Listening requests on specified port
	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
