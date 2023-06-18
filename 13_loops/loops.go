// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"time"
)

// Go has only the "for" loop structure

// Main function to be executed
func main() {
	fmt.Println("Loops")

	// Defining the counter variable
	i := 0

	// Creating a simple for loop
	for i < 10 {
		// Waiting for one second before continuing
		time.Sleep(time.Second * 1)
		fmt.Println("Adding up \"i\"", i)
		i++
	}

	fmt.Println("Final \"i\" value", i)

	// We can also initialize the counter variable for the "loop" structure
	for j := 0; j < 10; j++ {
		fmt.Println("Adding up \"j\"", j)
	}

	// We can use custom steps for the loop
	for j := 0; j < 10; j += 2 {
		fmt.Println("Adding up \"j\"", j)
	}

	// We can also iterate through items
	names := [3]string{"John", "Luke", "David"}

	// Loop to iterate through the names list
	for idx, name := range names {
		fmt.Println(idx, name)
	}

	// If we don't want to use the index
	for _, name := range names {
		fmt.Println(name)
	}

	// Loop to iterate through a string
	// This will return the ASCII code for the characters
	for idx, letter := range "Wordly Word" {
		// If we want to get the char for the ASCII code, we use 'string(char)'
		fmt.Println(idx, letter, string(letter))
	}

	// Creting a new map
	user := map[string]string{
		"name":    "Jack",
		"surname": "Henry",
	}

	// Loop to iterate through a map
	// In this case, it returns the key-value pair
	for key, value := range user {
		fmt.Println(key, value)
	}

	// Range for structs are not available

	// We can have infinite loops, like so
	// for {
	//		fmt.Println("Executing forever")
	//		time.Sleep(time.Second)
	// }
}
