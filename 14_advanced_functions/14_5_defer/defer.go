// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// With the 'defer' option, statements can be executed at the last possible moment
// Useful, for example, for closing database connections

func function1() {
	fmt.Println("Executing function 1")
}

func function2() {
	fmt.Println("Executing function 2")
}

func isStudentApproved(n1, n2 float32) bool {
	// Text to be displayed after the calculation is finished
	defer fmt.Println("Average value calculated. Results will be returned")

	fmt.Println("Function to check if student is approved")

	// Calculating the average value
	avg := (n1 + n2) / 2

	if avg >= 6 {
		return true
	}
	return false
}

// Main function to be executed
func main() {
	fmt.Println("Defer")

	// Here, we ask for the function 1 execution to be deferred until the last possible moment
	defer function1()
	// In this case, the function 2 will be executed before the function 1
	function2()
	fmt.Println(isStudentApproved(7, 8))

}
