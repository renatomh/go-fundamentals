// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Main function to be executed
func main() {
	fmt.Println("Maps")

	// Creating a new map
	// We define both the key and value types
	// All keys must be of the same type, as well as all the values
	user1 := map[string]string{
		"name":    "Peter",
		"surname": "Johnson",
	}

	fmt.Println(user1)

	// If we want to access a key-value from the maps, we must do like so
	fmt.Println(user1["name"], user1["surname"])

	// We can also have nested maps
	user2 := map[string]map[string]string{
		"name": {
			"first": "John",
			"last":  "Wick",
		},
		"course": {
			"name":   "Engineering",
			"campus": "Campus 1",
		},
	}

	fmt.Println(user2)
	// If we want to access a key-value from the maps, we must do like so
	fmt.Println(user2["name"]["first"], user2["name"]["last"], user2["course"]["name"], user2["course"]["campus"])

	// If we want to remove a key from a map
	delete(user2, "name")
	fmt.Println(user2)

	// If we want to add a new key to the map
	user2["address"] = map[string]string{
		"street": "John St.",
		"number": "324",
	}
	fmt.Println(user2)

}
