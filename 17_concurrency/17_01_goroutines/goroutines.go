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
	 - Time Frame: [t00 t01 t02 t03 t04 t05 t06 t07 t08 t09 t10]
	 - Core 1:     [ 1   2   1   1   1   2   3   2   1   2   1 ]

	* Parallelism is when tasks literally run at the same time, e.g., on a multicore processor.
	 - Time Frame: [t00 t01 t02 t03 t04 t05 t06 t07 t08 t09 t10]
	 - Core 1:     [ 1   1   1   1   1   1   1   1   1   1   1 ]
	 - Core 2:     [ 2   2   2   2               2   2         ]
	 - Core 3:     [             3   3   3   3   3             ]
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
