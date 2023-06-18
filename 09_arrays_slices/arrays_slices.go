// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"reflect"
)

// Arrays are lists with fixed-length in Go
// https://go.dev/tour/moretypes/6
// Slices are dinamically-sized lists
// https://go.dev/tour/moretypes/7

// Main function to be executed
func main() {
	fmt.Println("Arrays and Slices")

	// Creating a new array, we must provide the array length and type
	var array1 [5]string
	// Populating the array
	array1[0] = "1. Position 1"
	array1[1] = "1. Position 2"
	array1[2] = "1. Position 3"
	array1[3] = "1. Position 4"
	array1[4] = "1. Position 5"
	fmt.Println(array1)

	array2 := [5]string{"2. Position 1", "2. Position 2", "2. Position 3", "2. Position 4", "2. Position 5"}
	fmt.Println(array2)

	// Here, we can set the array size, according to the number of elements in the initialization
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	// Creating slices (which don't have fixed sizes)
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice1)

	// Showing different types for slices and arrays
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	// If we want to add new items to slices, we can append them
	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	// We can instantiate slices from arrays (defining the indexes)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Changing array elements by index will reflect on the slices (works like pointers)
	array2[1] = "2. New Position 2"
	fmt.Println(slice2)

	fmt.Println("----------")
	// Internal arrays
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Length
	fmt.Println(cap(slice3)) // Capacity

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Length
	fmt.Println(cap(slice3)) // Capacity

	// When Go identifies that the slice capacity will be exceeded
	// it doubles its capacity according the the current length
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Length
	fmt.Println(cap(slice3)) // Capacity

	// If we don't provide the slice capacity, it will be the length, by default
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Length
	fmt.Println(cap(slice4)) // Capacity

	// Once we append a new item, it will double the capacity
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Length
	fmt.Println(cap(slice4)) // Capacity

}
