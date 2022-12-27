// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Defining new structs
type user struct {
	name string
	age  uint8
	// We can have structs inside structs
	address address
}

type address struct {
	street string
	number uint8
}

// Main function to be executed
func main() {
	fmt.Println("Structs file")

	// Creating a new struct instance
	var user1 user
	// Filling struct fields
	user1.name = "Renato"
	user1.age = 28
	fmt.Println(user1)

	// Other ways to create a new instance
	address1 := address{"John St.", 8}

	user2 := user{"John", 36, address1}
	fmt.Println(user2)

	user3 := user{name: "David"}
	fmt.Println(user3)

	user4 := user{age: 26}
	fmt.Println(user4)
}
