// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Simple function (passing only a copy of the variable value as the parameter)
func signInversion(number int) int {
	return number * -1
}

// Function with pointer parameter (passing a reference for the variable as the parameter)
func signInversionWithPointer(number *int) int {
	// Here, we'll set the result as the number memory address content
	*number = *number * -1
	return *number
}

// Main function to be executed
func main() {
	fmt.Println("Functions with pointers")

	// Without using pointers on function parameters
	number := 20
	fmt.Println(number)
	invertedNumber := signInversion(number)
	fmt.Println(number, invertedNumber)

	// Using pointers on function parameters
	newNumber := 40
	fmt.Println(newNumber)
	invertedNewNumber := signInversionWithPointer(&newNumber)
	fmt.Println(newNumber, invertedNewNumber)

}
