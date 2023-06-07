// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Functions with more than one return and named returns
func mathCalculations(n1, n2 int8) (sum int8, subtraction int8) {
	// We don't need to define a variable, since its already defined to be returned by the name
	sum = n1 + n2
	subtraction = n1 - n2
	// Here, we'll return all the name return items
	return
}

// Main function to be executed
func main() {
	// Calling functions with more than one return
	mathResultsSum1, mathResultsSubtraction1 := mathCalculations(10, 15)
	fmt.Println(mathResultsSum1, mathResultsSubtraction1)
	// Retrieving only one value
	_, mathResultsSubtraction2 := mathCalculations(10, 15)
	fmt.Println(mathResultsSubtraction2)

}
