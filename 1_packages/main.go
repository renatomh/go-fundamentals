// Only "main" packages can be executed
package main

// Defining imports
import (
	"firstmod/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

// Main function to be executed
func main() {
	fmt.Println("Hello, world!")
	// Calling auxiliar function
	auxiliar.Write()

	// Checking for an email format
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
