package main

import "fmt"

// Reference: https://www.educative.io/answers/how-to-implement-a-stack-in-golang

// Defining the stack struct
type Stack []int

// Method to push a new item to the stack
func (s *Stack) Push(item int) {
	// Here, we simply add the item to the top of the stack
	*s = append(*s, item)
}

// Method to pop an item from the stack
func (s *Stack) Pop() int {
	// First we check if the stack is not empty
	if s.IsEmpty() {
		// If it's, we return -1
		return -1
	}

	// Otherwise, we remove the last item from the stack and return it
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}

// Method to check if stack is empty
func (s Stack) IsEmpty() bool {
	// If it has no elements, we return true, otherwise, returns false
	return len(s) == 0
}

func main() {
	// Initial info printing
	fmt.Println("Stack (LIFO) implementation with go")

	// Initializing a new stack
	stack := Stack{}

	// Adding values to the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	// Showing the created stack
	fmt.Println("Initial stack: ")
	fmt.Println(stack)

	// Removing some items and showing stack after removal
	fmt.Println("Item popped from stack:", stack.Pop())
	fmt.Println("Item popped from stack:", stack.Pop())

	// Showing updated stack
	fmt.Println("Updated stack: ")
	fmt.Println(stack)
}
