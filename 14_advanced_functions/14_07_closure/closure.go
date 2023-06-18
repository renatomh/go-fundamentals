// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Closure functions will 'remember' variables from its source/origin
func closure() func() {
	text := "Inside 'closure' function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

// Main function to be executed
func main() {
	fmt.Println("Closure")

	text := "Inside 'main' function"
	fmt.Println(text)

	// For example, even though the 'text' variable was defined at 'main' and 'closure', when calling
	// the closure function, it'll remember its own 'text' variable
	newFunction := closure()
	newFunction()

}
