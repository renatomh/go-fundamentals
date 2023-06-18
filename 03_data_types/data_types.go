// Only "main" packages can be executed
package main

// Defining imports
import (
	"errors"
	"fmt"
)

// Main function to be executed
func main() {
	// Number data types
	// Integer values:
	// * int8 (up to 8 bits)
	// * int16 (up to 16 bits)
	// * int32 (up to 32 bits)
	// * int64 (up to 64 bits)
	var intNumber16 int16 = 100
	fmt.Println(intNumber16)

	// If we specify only int, it'll use the computer's default architecture (32 bits, 64 bits, etc.)
	var pcIntNumber1 int = 1000000
	pcIntNumber2 := -2000000
	fmt.Println(pcIntNumber1, pcIntNumber2)

	// Unsigned int values (uint)
	var uintNumber1 uint32 = 2000000000
	fmt.Println(uintNumber1)

	// Aliases
	// INT32 = RUNE
	var intNumber32 rune = 2147483647
	fmt.Println(intNumber32)
	// UINT8 = BYTE
	var uintNumber2 byte = 255
	fmt.Println(uintNumber2)

	// Float values:
	// * float32 (up to 32 bits)
	// * float64 (up to 64 bits)
	var floatNumber1 float32 = 123.45
	fmt.Println(floatNumber1)

	var floatNumber2 float64 = 12300000.45
	fmt.Println(floatNumber2)

	floatNumber3 := -12345.67
	fmt.Println(floatNumber3)

	// String Data Types
	var str1 string = "Text Value 1"
	fmt.Println(str1)

	str2 := "Text Value 2"
	fmt.Println(str2)

	// If we want to get the ASCII value for a character
	// We use single quotes ''
	asciiCharNumber := 'B'
	fmt.Println(asciiCharNumber)

	// Zero-value: assigned to a variable when we don't initialize it
	var zeroStr string
	fmt.Println(zeroStr)

	var zeroInt int16
	fmt.Println(zeroInt)

	// Boolean values
	var bool1 bool = true
	fmt.Println(bool1)
	bool2 := false
	fmt.Println(bool2)
	var bool3 bool
	fmt.Println(bool3)

	// Error values
	var error1 error
	fmt.Println(error1)

	var error2 error = errors.New("Internal Error")
	fmt.Println(error2)

}
