// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"time"
)

// Channels allows sending and receiving data

// Main function to be executed
func main() {
	fmt.Println("Channels")

	// Creating a new channel variable
	channel := make(chan string)

	// Calling the function with the channel
	go write("Just keep swimming", channel)

	// Giving feedback to the user
	fmt.Println("After 'write' function execution starts")

	// One way to deal with the channel
	/*
		for {
			// Waiting for the channel to receive the value and checking if it's open
			message, open := <-channel
			// If channel is not open, we exit the loop
			if !open {
				break // Exist the loop
			}
			// Showing the received message
			fmt.Println(message)
		}
	*/

	// Another way to deal with channel usage
	// This already checks for open/closed channels
	for message := range channel {
		// Showing the received message
		fmt.Println(message)
	}

	fmt.Println("Tasks done")

}

// This function will receive some data and return to the channel
func write(text string, channel chan string) {
	// Creating an fixed size loop
	for i := 0; i < 5; i++ {
		// Returning data to the channel
		channel <- text
		time.Sleep(time.Second)
	}

	// Closing the channel after function finishes
	close(channel)
}
