// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Function to return the week day from number
func weekDay(number int) string {
	// Creating the switch structure
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid Number"
	}
}

// Alternative function to return the week day from number
func altWeekDay(number int) string {
	// Initializing value to be returend
	var weekDay string

	// Creating the switch structure
	switch {
	case number == 1:
		weekDay = "Sunday"
		// The 'fallthrough' clause makes the next case to be executed
		// in this case, when number == 1, the 'weekDay' would be set as "Monday"
		// fallthrough
	case number == 2:
		weekDay = "Monday"
	case number == 3:
		weekDay = "Tuesday"
	case number == 4:
		weekDay = "Wednesday"
	case number == 5:
		weekDay = "Thursday"
	case number == 6:
		weekDay = "Friday"
	case number == 7:
		weekDay = "Saturday"
	default:
		weekDay = "Invalid Number"
	}

	// Returning obtained value
	return weekDay
}

// Main function to be executed
func main() {
	fmt.Println("Switch")

	// Calling the function with switch statement
	day := weekDay(3)
	fmt.Println(day)
	day = weekDay(6)
	fmt.Println(day)
	day = weekDay(8)
	fmt.Println(day)

	// Calling the alternate function with switch statement
	day = altWeekDay(3)
	fmt.Println(day)
	day = altWeekDay(6)
	fmt.Println(day)
	day = altWeekDay(8)
	fmt.Println(day)

}
