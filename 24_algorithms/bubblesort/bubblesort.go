package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Reference: https://blog.betrybe.com/tecnologia/bubble-sort-tudo-sobre/

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

// Bubble sort function implementation (sorts the slice in-place)
func bubblesort(list []int) []int {
	// For this algorithm, we compare every pair of values in the list
	// It might need multiple runs to sort all the list

	// Setting a flag to check if slice is sorted
	swapped := true

	// While slice is still not sorted (that is, swaps had to be made)
	for swapped == true {
		// Resetting the sorting flag
		swapped = false
		// Check all elements in the slice
		for i := 0; i < len(list)-1; i++ {
			// If previous element is greater than the next one
			if list[i] > list[i+1] {
				// We'll swap elements
				list[i], list[i+1] = list[i+1], list[i]
				// And update the sorting flag, since elements had to be swapped
				swapped = true
			}
		}
	}

	// Finally, we return the sorted slice
	return list
}

func main() {
	// Initial info printing
	fmt.Println("Bubble sort algorithm implementation with go")

	// Here we can either create a new slice/list manually, or generate a random one
	manSlice := []int{22, 4, 3, 7, 98, 4, 5, 6, 12, 3, 6, 11, 10, 13, 19, 28, 22, 1, 5, 7, 8, 9}
	randomSlice := generateRandomSlice(35, 50)

	// Presenting the original slices, sorting and showing the resulting slices
	fmt.Println("Manually created slice:")
	fmt.Println("> Original:\t", manSlice)
	bubblesort(manSlice)
	fmt.Println("> Sorted:\t", manSlice)
	fmt.Println("Random slice:")
	fmt.Println("> Original:\t", randomSlice)
	bubblesort(randomSlice)
	fmt.Println("> Sorted:\t", randomSlice)
}
