package main

import (
	"fmt"
	"time"
)

// Generators encapsulate calls on a goroutine and return it
// They're used to 'hide the complexity' of goroutine functions

func main() {
	fmt.Println("Generators")

	// Calling the function which encapsulates the calls on the goroutine
	// This channel will only receive data
	channel := write("Just keep swimming")

	// Loop to read the channel data
	for i := 0; i < 10; i++ {
		// Printing data which is received by the channel
		fmt.Println(<-channel)
	}

}

// Creating function to write data, setting return type as a one-way channel
func write(text string) <-chan string {
	// Initializing the channel
	channel := make(chan string)

	// Creating the goroutine function
	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Returning the created channel
	return channel
}
