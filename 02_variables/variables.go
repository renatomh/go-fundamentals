// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Main function to be executed
func main() {
	// Defining variables with explicit type
	var strvar1 string = "String Variable 1"
	fmt.Println(strvar1)
	// Defining variables with implicit type
	strvar2 := "String Variable 2"
	fmt.Println(strvar2)

	// Declaring multiple variable at once
	var (
		strvar3 string = "String Variable 3"
		intvar1 int    = 12
	)
	fmt.Println(strvar3, intvar1)

	// Declaring multilpe variable of the same type at once
	strvar4, strvar5 := "String Variable 4", "String Variable 5"
	fmt.Println(strvar4, strvar5)

	// Declaring constant values
	const strconst1 string = "String Const 1"
	fmt.Println(strconst1)

	// If we want to exchange two variables values, we don't need an aux variable
	// we can do it like so
	strvar4, strvar5 = strvar5, strvar4
	fmt.Println(strvar4, strvar5)

}
