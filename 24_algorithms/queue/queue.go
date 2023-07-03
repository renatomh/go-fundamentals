package main

import "fmt"

// Reference: https://www.geeksforgeeks.org/queue-in-go-language/

// Defining the queue struct
type Queue []int

// Method to enqueue a new item to the queue
func (q *Queue) Enqueue(item int) {
	// Here, we simply add the item to the end of the queue
	*q = append(*q, item)
}

// Method to dequeue an item from the queue
func (q *Queue) Dequeue() int {
	// First we check if the queue is not empty
	if q.IsEmpty() {
		// If it's, we return -1
		return -1
	}

	// Otherwise, we remove the first item from the queue and return it
	dequeue := (*q)[0]
	*q = (*q)[1:]
	return dequeue
}

// Method to check if queue is empty
func (q Queue) IsEmpty() bool {
	// If it has no elements, we return true, otherwise, returns false
	return len(q) == 0
}

func main() {
	// Initial info printing
	fmt.Println("Queue (FIFO) implementation with go")

	// Initializing a new queue
	queue := Queue{}

	// Adding values to the queue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)

	// Showing the created queue
	fmt.Println("Initial queue: ")
	fmt.Println(queue)

	// Removing some items and showing queue after removal
	fmt.Println("Item removed from queue:", queue.Dequeue())
	fmt.Println("Item removed from queue:", queue.Dequeue())

	// Showing updated queue
	fmt.Println("Updated queue: ")
	fmt.Println(queue)
}
