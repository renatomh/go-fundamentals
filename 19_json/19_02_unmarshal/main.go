package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// The Unmarshal function converts JSON to Go objects

type dog struct {
	// We can define the JSON keys on the struct fields
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

type namelessDog struct {
	// We can ignore JSON keys like below
	Name  string `json:"-"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	// Creating a new dog formatted as JSON
	dogJSON := `{"name":"Rex","breed":"Dalmata","age":3}`

	// Creating a new var to receive the dog data
	var d dog
	fmt.Println(d)
	// The Unmarshal function receives the JSON data (in bytes) and the memory address where we want to store the result
	if err := json.Unmarshal([]byte(dogJSON), &d); err != nil {
		log.Fatal(err)
	}
	// Showing the result from the Unmarshal function
	fmt.Println(d)

	// Ignoring JSON keys
	nDogJSON := `{"name":"Ugly Name","breed":"Husky","age":5}`
	var nDog namelessDog
	if err := json.Unmarshal([]byte(nDogJSON), &nDog); err != nil {
		log.Fatal(err)
	}
	fmt.Println(nDog)

	// Using maps
	dog2JSON := `{"name":"Toby","breed":"Poodle","age":"6"}`

	// Creating a new var to receive the dog data
	d2 := make(map[string]string)
	// Converting JSON to map
	if err := json.Unmarshal([]byte(dog2JSON), &d2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d2)

	// Here, all the JSON values must respect the map type (in this case, should be strings)
	// For example, the following object would lead to an error, since we are using an int for the age
	errorDogJSON := `{"name":"Max","breed":"Golden","age":12}`
	errorDog := make(map[string]string)
	// Converting JSON to map
	if err := json.Unmarshal([]byte(errorDogJSON), &errorDog); err != nil {
		log.Fatal(err)
	}
	fmt.Println(errorDog)

}
