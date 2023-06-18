// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
)

// Channels with buffer will only block after number of received messages reach the buffer size

// Main function to be executed
func main() {
	fmt.Println("Channels with Buffer")

	// When creating new channels, you can also set a buffer, to define how many items can be received
	// before locking the channel until a variable receives the channel item
	channel := make(chan string, 4) // In this, we set the channel to buffer up to 4 items

	// Sending the messages to the channel
	channel <- "Just keep swimming!"
	channel <- "Just keep swimming!!"
	channel <- "Still swimming!!!"
	channel <- "Escape!"
	// If we received more items and didn't output it, the code would enter a deadlock state
	// 4 + 1 = 5; 5 > 4 (channel buffer size)

	// Receiving the messages through the channel
	message1 := <-channel
	message2 := <-channel
	message3 := <-channel
	message4 := <-channel

	// Giving feedback to the user
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(message4)

	// After that, the buffer will be empty (since all messages were received)
	// and we can send more data to the channel
	channel <- "Haven't we escaped yet?"
	channel <- "Now we have!!"

	// Receiving messages and printing
	message1 = <-channel
	message2 = <-channel
	fmt.Println(message1)
	fmt.Println(message2)
}
