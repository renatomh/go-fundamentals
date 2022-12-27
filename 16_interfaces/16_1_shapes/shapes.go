// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"math"
)

// Creating the shape interface, which will have common fields
type shape interface {
	area() float64
}

// Function to write the shape of an area
func writeArea(s shape) {
	fmt.Printf("The shape area is %0.2f\n", s.area())
}

// Creating the rectangle struct
type rectangle struct {
	height float64
	width  float64
}

// Method to calculate the rectangle area
func (r rectangle) area() float64 {
	return r.height * r.width
}

// Creating the circle struct
type circle struct {
	radius float64
}

// Method to calculate the circle area
func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
	// We could also use the 'math.Pow()' function
	// return math.Pi * math.Pow(c.radius, 2)
}

// Main function to be executed
func main() {
	fmt.Println("Interfaces")

	// Creating a rectangle shape
	r := rectangle{10, 15}
	// Calling the function to display the rectangle area
	writeArea(r)

	// Creating a circle shape
	c := circle{10}
	// Calling the function to display the circle area
	writeArea(c)
}
