package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// The Marshal function converts Go objects to JSON

type dog struct {
	// We can define the JSON keys on the struct fields
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	// Creating a new dog struct
	d := dog{"Rex", "Dalmata", 3}
	fmt.Println(d)

	// Converting the struct to a JSON bytes object
	dogJSON, error := json.Marshal(d)

	// If an error occurs
	if error != nil {
		log.Fatal(error)
	}

	// Showing the raw bytes return of JSON Marshal function
	fmt.Println(dogJSON)
	// Printing as a JSON string
	fmt.Println(bytes.NewBuffer(dogJSON))
	// We could also do it like this
	fmt.Println(string(dogJSON))

	// If we were using maps
	d2 := map[string]string{
		"name":  "Toby",
		"breed": "Poodle",
		"age":   "6",
	}
	dog2JSON, err := json.Marshal(d2)

	// If an error occurs
	if err != nil {
		log.Fatal(err)
	}
	// Showing the raw bytes return of JSON Marshal function
	fmt.Println(dog2JSON)
	// Printing as a JSON string
	fmt.Println(string(dog2JSON))

}
