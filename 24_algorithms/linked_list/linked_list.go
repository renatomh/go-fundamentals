package main

import "fmt"

// Reference: https://www.golangprograms.com/golang-program-for-implementation-of-linked-list.html

// Defining the Node struct
type Node struct {
	data int
	next *Node
}

// Defining the list with the head (first Node) struct
type List struct {
	head *Node
}

// Method to add a new Node to the the end list
func (l *List) add(value int) {
	// Creating the new Node to be added, with provided data and null next Node
	newNode := &Node{data: value}

	// If the list does not have a head (first node yet)
	if l.head == nil {
		// We simply set the new Node as the head and return
		l.head = newNode
		return
	}

	// Now, if there's a head set, we run through the head (first node) until the tail (last node)
	currNode := l.head
	for currNode.next != nil {
		currNode = currNode.next
	}

	// And add the node as the list tail (last node)
	currNode.next = newNode
}

// Method to remove an existing Node from the list
func (l *List) remove(value int) {
	// If the list is empty, we simply return
	if l.head == nil {
		return
	}

	// Otherwise, if the value is in the head (first Node)
	if l.head.data == value {
		// We set the head as the next node and return
		l.head = l.head.next
		return
	}

	// Else, we'll run through the linked list until we find the node with the value
	currNode := l.head
	for currNode.next != nil && currNode.next.data != value {
		currNode = currNode.next
	}

	// If the value was found, we'll update the previous node to use the next one as its own next
	if currNode.next != nil {
		currNode.next = currNode.next.next
	}
}

// Function to print an existing linked list
func printList(l *List) {
	// Initializing at the head
	currNode := l.head
	// Running through all nodes
	for currNode != nil {
		// Printing the node data
		fmt.Printf("%d ", currNode.data)
		// And going to the next node
		currNode = currNode.next
	}
	fmt.Println()
}

func main() {
	// Initial info printing
	fmt.Println("Linked lists implementation with go")

	// Initializing a new linked list
	list := &List{}

	// Adding values to the list
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)
	list.add(6)

	// Showing the created list
	fmt.Println("Initial List: ")
	printList(list)

	// Removing some items and showing list after removal
	list.remove(2)
	fmt.Println("List after removing 2: ")
	printList(list)
	list.remove(4)
	fmt.Println("List after removing 4: ")
	printList(list)
}
