// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Methods are necessarily attached to a struct

// Defining a 'user' struct
type user struct {
	name string
	age  uint8
}

// Creating a method for the 'user' struct to save user's data
func (u user) save() {
	fmt.Printf("Saving '%s' user's data on database\n", u.name)
}

// Creating a method for the 'user' struct to check if user's age is over 18
func (u user) ofLegalAge() bool {
	return u.age >= 18
}

// If we want to change a user's field inside a method, we can use pointers
func (u *user) haveBirthday() {
	u.age++
}

// Main function to be executed
func main() {
	// Creating new users instances
	user1 := user{"John Cage", 18}
	user2 := user{"Matt Damon", 37}
	user3 := user{"Michael James", 16}

	// Calling the user's methods
	// To save the user
	user1.save()
	user2.save()

	// Check if user is of legal age
	if user2.ofLegalAge() {
		fmt.Printf("%s is of legal age\n", user2.name)
	} else {
		fmt.Printf("%s is not of legal age\n", user2.name)
	}
	if user3.ofLegalAge() {
		fmt.Printf("%s is of legal age\n", user3.name)
	} else {
		fmt.Printf("%s is not of legal age\n", user3.name)
	}

	// Add one year to the user's age
	fmt.Println(user1.age)
	user1.haveBirthday()
	fmt.Println(user1.age)
	fmt.Println(user3.age)
	user3.haveBirthday()
	fmt.Println(user3.age)

}
