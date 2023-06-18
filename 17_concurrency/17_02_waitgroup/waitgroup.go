// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"sync"
	"time"
)

// With WaitGroups, we can sync the execution of goroutines, such that once they're all done
// we'll consider the main execution is completed

// Main function to be executed
func main() {
	fmt.Println("WaitGroup")

	// In this scenario, the first text will be written 5x, then, the second one for 5x as well
	//write("Just keep swimming")
	//write("Escape!")

	// Creating a waitGroup variable
	var waitGroup sync.WaitGroup

	// Specifying number of goroutines which will be executed
	waitGroup.Add(2)

	// Creating anonymous functions with goroutines
	go func() {
		write("Just keep swimming")
		// When it's over, we call the 'done' function from the waitGroup
		waitGroup.Done() // Removes 1 from the waitGroup tasks quantity
	}()
	go func() {
		write("Escape!")
		// When it's over, we call the 'done' function from the waitGroup
		waitGroup.Done() // Removes 1 from the waitGroup tasks quantity
	}()

	// Asking the waitGroup execution to run all tasks (reach zero)
	// before cotinuing the code
	waitGroup.Wait()

	fmt.Println("Tasks done")

}

// This function will run for a fixed number of times
func write(text string) {
	// Creating an fixed size loop
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
