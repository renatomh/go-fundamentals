// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Variadic functions can receive different numbers of parameters
func sum(numbers ...int) int {
	// Creating the total variable to be returned
	total := 0
	// Iterating through all provided numbers
	for _, number := range numbers {
		total += number
	}
	// Returning the total value
	return total
}

// Variadic function which receives both string and number parameters
// Only one type can receive multiple parameters, and must be placed in the end
func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

// Main function to be executed
func main() {
	// Calling the function with different numbers of parameters
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(5, 10, 15))
	fmt.Println(sum())

	// Calling second variadic function
	write("See the number", 1, 2, 3, 4, 5)

}
