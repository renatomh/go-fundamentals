// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Main function to be executed
func main() {
	fmt.Println("Channels with Buffer")

	// When creating new channels, you can also set a buffer, to define how many items can be received
	// before locking the channel until a variable receives the channel item
	channel := make(chan string, 4) // In this, we set the channel to buffer up to 4 items

	// Sending the messages to the channel
	channel <- "Just keep swimming!"
	channel <- "Just keep swimming!!"
	// If we received 3 more items and didn't output it, the code would enter a deadlock state
	// 2 + 3 = 5; 5 > 4 (channel buffer size)

	// Receiving the messages through the channel
	message1 := <-channel
	message2 := <-channel

	// Giving feedback to the user
	fmt.Println(message1)
	fmt.Println(message2)
}
