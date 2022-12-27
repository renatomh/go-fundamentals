// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Defining new structs
type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	// This is similar to "hierarchy" in object oriented languages
	person
	course  string
	college string
}

// Main function to be executed
func main() {
	fmt.Println("Hierarchy")

	// Creating new struct instances
	person1 := person{"John", "Wick", 27, 180}
	fmt.Println(person1)

	student1 := student{person1, "Engineering", "MIT"}
	fmt.Println(student1)

	student2 := student{person{"Jack", "Ryan", 31, 184}, "IT", "Oxford"}
	fmt.Println(student2)

	// We can access the student's person attributes directly
	fmt.Println(student1.name, student1.age, student1.height)
}
