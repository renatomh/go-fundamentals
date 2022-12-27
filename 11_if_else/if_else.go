// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Main function to be executed
func main() {
	fmt.Println("Control structures")

	number1 := 10

	// Making comparisons
	if number1 > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Smaller than or equal to 15")
	}

	// We can initialize a variable during an 'if' condition
	// However, we won't be able to access it outside the 'if' scope
	if number2 := number1; number2 > 0 {
		fmt.Println("Greater than 0")
	} else if number2 == 0 {
		// We can also have 'else if' conditions
		fmt.Println("Equal to 0")
	} else {
		fmt.Println("Smaller than 0")
	}

}
