// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Pointers are variables which holds the memory address of other variables
// https://go.dev/tour/moretypes/1

// Main function to be executed
func main() {
	fmt.Println("Pointers")

	// Creating a new variable
	var var1 int = 10
	// Here, we simply copy the current var1 value to var2
	var var2 int = var1

	// If we add the var1 value, var2 will remain the same
	fmt.Println(var1, var2)
	var1++
	fmt.Println(var1, var2)
	var2--
	fmt.Println(var1, var2)

	// Pointers refers to memory addresses
	var var3 int
	var pointer *int

	var3 = 100
	// Here, we save the var3 memory address to the pointer
	pointer = &var3
	// Showing the variable 3, pointer content (var 3 mem address) and mem address content (variable 3 value)
	// This '*' before the pointer variable is known as "dereferencing" or "indirecting".
	fmt.Println(var3, pointer, *pointer)

	// If we change the variable 3 value, mem address content will also change
	// So these changes will be reflected on the pointer
	var3 = 150
	fmt.Println(var3, pointer, *pointer)

}
