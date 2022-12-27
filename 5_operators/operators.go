// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Main function to be executed
func main() {
	// Arithmetics operators: +, -, /, *, %
	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	mult := 10 * 5
	rem := 10 % 4

	fmt.Println(sum, subtraction, division, mult, rem)

	// Operations can only be done with variables from the same type
	// The following would lead to an error
	// var number1 int16 = 10
	// var number2 int32 = 25
	// sumValue := number1 + number2

	// Assignment operators
	var var1 string = "String 1"
	var2 := "String 2"

	fmt.Println(var1, var2)

	// Comparison operators
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Logical operators
	trueValue, falseValue := true, false
	fmt.Println("True and False:", trueValue && falseValue)
	fmt.Println("True or False:", trueValue || falseValue)
	fmt.Println("Not True:", !trueValue)

	// Unary operators
	number := 10
	fmt.Println("Original Number:", number)

	number++
	number += 15
	fmt.Println("Added Number:", number)

	number--
	number -= 20
	fmt.Println("Subtracted Number:", number)

	number *= 2
	fmt.Println("Multiplied Number:", number)

	number /= 3
	fmt.Println("Divided Number:", number)

	number %= 2
	fmt.Println("Remainder Number:", number)
}
