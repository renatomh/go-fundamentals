package main

import "fmt"

// Worker pools are like tasks queues
// Useful when you have a large queue of tasks to be executed

func main() {
	fmt.Println("Worker Pools")

	// Defining the position for the desired Fibonnaci number
	fibonacci_position := 45

	// Creating channels for the tasks to be executed and returned results
	// The buffer size will be defined as the same as the position for the desired Fibonnaci number
	tasks := make(chan int, fibonacci_position)
	results := make(chan int, fibonacci_position)

	// Calling the workers with 'goroutines'
	// The more workers we use, the better, however, we're limited to our machine's computational resources
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	// Loop to fill tasks channel with data
	for i := 0; i < fibonacci_position; i++ {
		tasks <- i
	}
	// Closing the tasks channel, so it won't receive any more data
	close(tasks)

	// Showing the results for the user
	for i := 0; i < fibonacci_position; i++ {
		// Whenver the result is sent to the channel, it's received here
		result := <-results
		fmt.Println(result)
	}

}

// This function is used to execute the tasks defined on the channel
// We also set the 'tasks' channel to only receive data and the 'results' channel to only send data
func worker(tasks <-chan int, results chan<- int) {
	// For each number received in the tasks channel
	for number := range tasks {
		// We set the resulting fibonnaci to the 'results' channel
		results <- fibonacci(number)
	}
}

// Recursive function to calculate the Fibonacci numbers at a specific position
func fibonacci(position int) int {
	// We must define a stop condition, in order to avoid stack overflows
	if position <= 1 {
		return position
	}

	// Recursive statement
	return fibonacci(position-2) + fibonacci(position-1)
}
