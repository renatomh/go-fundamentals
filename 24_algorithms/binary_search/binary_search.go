package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Reference: https://pkg.go.dev/sort#Ints

// Here we have a function to generate a slice (list) with random values
func generateRandomSlice(size int, limit int) []int {
	// First, we initialize the slice
	slice := make([]int, size)

	// Now, we'll seed the random numbers generation with current time
	rand.Seed(time.Now().UnixNano())

	// Adding random elements to the slice
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(limit)
	}

	// Returning the generated slice
	return slice
}

// Binary search function implementation (returns the number index or -1 if not found)
// It only works with sorted lists
func binarySearch(lookFor int, list []int) int {
	// If the slice is empty, no nmber can be found
	if len(list) == 0 {
		return -1
	}

	// Defining an initial index for the number to be found
	index := 0

	// Now, we'll split the slice in half, until there's only one item left
	for len(list) > 1 {
		// Splitting the slice in two halfs by getting the middle index
		middleIndex := len(list) / 2
		lowerHalf := list[:middleIndex]
		upperHalf := list[middleIndex:]

		// Checking if the number we're searching is in the lower or upper half
		if lookFor < list[middleIndex] {
			// In this case, we'll get only the lower half
			list = lowerHalf
		} else {
			// Otherwise, the new list will be the upper half
			list = upperHalf
			// And we must also update the index, since we're ignoring the items on lower half
			index += len(lowerHalf)
		}

		// Showing the lower and upper halfs for the user (displaying how the search is going)
		fmt.Println("Lower half / Upper Half\n", lowerHalf, "[", lookFor, "]", upperHalf)
	}

	// Finally, if no number was found, we return the -1 index
	if lookFor != list[0] {
		return -1
	}

	// Otherwise, we'll return the corresponding index
	return index
}

func main() {
	// Initial info printing
	fmt.Println("Binary search algorithm implementation with go")

	// Generate a random slice and sorting it
	randomSlice := generateRandomSlice(75, 150)
	sort.Ints(randomSlice)

	// Defining a number to be found in the slice
	number := rand.Intn(150)

	// Presenting the original slice and wanted number
	fmt.Println("Random generated slice:", randomSlice)
	fmt.Println("Number to be found in the slice:", number)

	// Looking for the index of the desired number in the list/slice
	numberIndex := binarySearch(number, randomSlice)

	// Showing the final results
	if numberIndex != -1 {
		fmt.Printf("Index for the found number (%d) in the slice: %d", randomSlice[numberIndex], numberIndex)
	} else {
		fmt.Printf("Number (%d) could not be found in the slice.", number)
	}
}
