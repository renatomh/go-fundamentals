package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Reference: https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html

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

// Quicksort function implementation (sorts the slice in-place)
func quicksort(list []int) []int {
	// First, we check if the list has more than one element
	if len(list) < 2 {
		return list
	}

	// Here, we initialize the left and right indexes as the first and last positions
	left, right := 0, len(list)-1

	// We'll also pick a random pivot index to start the algorithm
	pivot := rand.Int() % len(list)

	// Now, we swap the elements at the pivot and right indexes
	list[pivot], list[right] = list[right], list[pivot]

	// And for each other element in the list, we move if it's smaller
	for i, _ := range list {
		// If element at current position is smaller than the right edge one
		if list[i] < list[right] {
			// We swap the elements
			list[left], list[i] = list[i], list[left]
			// And update the left index
			left++
		}
	}

	// Now, we swap the left and right indexes
	list[left], list[right] = list[right], list[left]

	// And recursively call the function for the splitted slices
	quicksort(list[:left])
	quicksort(list[left+1:])

	// Finally, we return the sorted slice
	return list
}

func main() {
	// Initial info printing
	fmt.Println("Quicksort algorithm implementation with go")

	// Here we can either create a new slice/list manually, or generate a random one
	manSlice := []int{22, 4, 3, 7, 98, 4, 5, 6, 12, 3, 6, 11, 10, 13, 19, 28, 22, 1, 5, 7, 8, 9}
	randomSlice := generateRandomSlice(35, 50)

	// Presenting the original slices, sorting and showing the resulting slices
	fmt.Println("Manually created slice:")
	fmt.Println("> Original:\t", manSlice)
	quicksort(manSlice)
	fmt.Println("> Sorted:\t", manSlice)
	fmt.Println("Random slice:")
	fmt.Println("> Original:\t", randomSlice)
	quicksort(randomSlice)
	fmt.Println("> Sorted:\t", randomSlice)
}
