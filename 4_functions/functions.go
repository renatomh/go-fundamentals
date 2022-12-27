// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Creating a sum function
func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Functions with more than one return
func mathCalculations(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

// Main function to be executed
func main() {
	// Calling the function to sum numbers
	sumValue := sum(10, 20)
	fmt.Println(sumValue)

	// Creating functions as variables
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fResult := f("Calling the variable function.")
	fmt.Println(fResult)

	// Calling functions with more than one return
	mathResultsSum1, mathResultsSubtraction1 := mathCalculations(10, 15)
	fmt.Println(mathResultsSum1, mathResultsSubtraction1)
	// Retrieving only one value
	_, mathResultsSubtraction2 := mathCalculations(10, 15)
	fmt.Println(mathResultsSubtraction2)

}
