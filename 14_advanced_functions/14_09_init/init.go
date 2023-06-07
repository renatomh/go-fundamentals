// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Defining a global variable
var n int

// The 'init' function will be executed before the 'main' function
// Even if it's placed after the 'main' function
// Each file in the package can have an 'init' function
// Useful for setting up variables and settings before 'main' function execution
func init() {
	fmt.Println("Executing 'init' function")
	n = 10
}

// Main function to be executed
func main() {
	fmt.Println("Executing 'main' function")
	fmt.Println(n)

}
