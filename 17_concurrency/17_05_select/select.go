// Only "main" packages can be executed
package main

// Defining imports
import (
	"fmt"
	"time"
)

// 'Selects' are similar to 'switches', however it'll test cases for channels outputs

// Main function to be executed
func main() {
	fmt.Println("Select Channels")

	// Creating new channels
	channel1, channel2 := make(chan string), make(chan string)

	// Setting (and calling) an anonymous function to send data for the first channel
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	// Setting (and calling) an anonymous function to send data for the second channel
	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	// If we didn't use selects, the channel which receives data faster would be limited by the
	// channel which receives data slower
	for i := 0; i < 3; i++ {
		channel1Message := <-channel1
		fmt.Println(channel1Message)
		channel2Message := <-channel2
		fmt.Println(channel2Message)
	}

	// Using selects, whichever channel is ready to output data will be able to do it
	// without having to wait for other channels output to be ready
	for {
		select {
		case channel1Message := <-channel1:
			fmt.Println(channel1Message)
		case channel2Message := <-channel2:
			fmt.Println(channel2Message)
		}
	}

}
