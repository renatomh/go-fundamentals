package auxiliar

// Defining imports
import (
	"fmt"
)

// Main function name must start with a capital letter to be exported
// Otherwise, it would only be visible inside this package

// Prints a message on the screen
func Write() {
	fmt.Println("Hello, from aux package!")
}
