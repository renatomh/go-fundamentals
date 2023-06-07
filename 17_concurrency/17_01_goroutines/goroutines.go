// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"time"
)

/*
	Concurrency != Parallelism

	* Concurrency is when two or more tasks can start, run, and complete in overlapping time periods.
	It doesn't necessarily mean they'll ever both be running at the same instant.
	For example, multitasking on a single-core machine.

	* Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.
*/

// Main function to be executed
func main() {
	fmt.Println("Concurrency")

	// In this scenario, the first text will be written forever
	//write("Just keep swimming")
	//write("Escape!")

	// In this case, using 'goroutine', it starts the first function execution
	// but it doesn't wait for it to finish, to keep running the code
	go write("Just keep swimming")
	write("Escape!")

}

// This function will run continuously
func write(text string) {
	// Creating an infinite loop
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
