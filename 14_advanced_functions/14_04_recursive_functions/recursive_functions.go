// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// An example of a recursive function is a function to calculate the Fibonacci numbers at a specific position
// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89...
func fibonacci(position uint) uint {
	// We must define a stop condition, in order to avoid stack overflows
	if position <= 1 {
		return position
	}

	// Recursive statement
	return fibonacci(position-2) + fibonacci(position-1)
}

// Main function to be executed
func main() {
	fmt.Println("Recursive functions")

	// Defining a position and calling the fibonacci function
	position := uint(15)
	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(uint(i)))
	}

	// It's better to avoid recursive functions when too many calculations/executions must be performed

}
