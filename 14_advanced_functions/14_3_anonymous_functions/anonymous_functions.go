// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Main function to be executed
func main() {
	// Creating and calling an anonymous function
	anonymousReturn := func(text string) string {
		return fmt.Sprintf("Received -> %s %d", text, 20)
	}("Anonymous parameter for anonymous function")

	fmt.Println(anonymousReturn)
}
