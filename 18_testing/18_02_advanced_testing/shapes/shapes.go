package shapes

// Defining imports
import (
	"math"
)

// Creating the shape interface, which will have common fields
type Shape interface {
	Area() float64
}

// Creating the rectangle struct
type Rectangle struct {
	Height float64
	Width  float64
}

// Method to calculate the rectangle area
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Creating the circle struct
type Circle struct {
	Radius float64
}

// Method to calculate the circle area
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
	// We could also use the 'math.Pow()' function
	// return math.Pi * math.Pow(c.radius, 2)
}
