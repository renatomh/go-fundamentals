// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// When a code execution 'panics', it'll call all the 'defer' statements
// There, we can try to recover the code execution
func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Code execution was successfully recovered!")
	}
}

func isStudentApproved(n1, n2 float32) bool {
	// Setting a defer to the recover method
	defer recoverExecution()

	// Calculating the average value
	avg := (n1 + n2) / 2

	// If average value is greater than or smaller than 6
	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	// In this case, if the average value is equal to 6, we'll throw a 'panic' statement
	// This will stop the code execution
	panic("THE AVERAGE VALUE IS EXACTLY 6")
}

// Main function to be executed
func main() {
	fmt.Println("Panic and Recover")

	fmt.Println(isStudentApproved(6, 7))
	fmt.Println("After execution")
	fmt.Println(isStudentApproved(5, 4))
	fmt.Println("After execution")
	fmt.Println(isStudentApproved(6, 6))
	fmt.Println("After execution")

}
