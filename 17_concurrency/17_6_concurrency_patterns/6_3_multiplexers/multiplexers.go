package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplexers are useful to unite different channels into a single channel, centralizing communication

func main() {
	fmt.Println("Multiplexers")

	// Calling the function which selects the input channels
	channel := mux(
		write("They're taking the hobbits to Isengard!"),
		write("What did you say?"),
		write("The Hobbits, The Hobbits, The Hobbits to Isengard, to Isengard!"))

	// Loop to read the channel data
	for i := 0; i < 15; i++ {
		// Printing data which is received by the channel
		fmt.Println(<-channel)
	}

}

// Creating the multiplexer from the different input channels, directing data to the
// output channel which will be returned by the function
func mux(inputChannel1, inputChannel2, inputChannel3 <-chan string) <-chan string {
	// Initizaling the output channel
	outputChannel := make(chan string)

	// Goroutine to
	go func() {
		for {
			// Whichever input channel is ready will be used to direct data to the output channel
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			case message := <-inputChannel3:
				outputChannel <- message
			}
		}
	}()

	// Returning the created output channel
	return outputChannel
}

// Creating function to write data, setting return type as a one-way channel
func write(text string) <-chan string {
	// Initializing the channel
	channel := make(chan string)

	// Creating the goroutine function
	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			// We'll use random durations (between 0 and 2000 ms)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	// Returning the created channel
	return channel
}
