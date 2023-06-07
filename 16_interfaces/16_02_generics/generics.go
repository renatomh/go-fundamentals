// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Generic interfaces are useful when a lot of flexibility is required
// For example, the fmt.Println() function, which can print anything in the screen

// Creating a function which receives a generic interface
func generic(interf interface{}) {
	fmt.Println(interf)
}

// Main function to be executed
func main() {
	fmt.Println("Generic Interfaces")

	generic("String")
	generic(1)
	generic(true)

	// Creating a map with generic interfaces
	mapItem := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapItem)

}
