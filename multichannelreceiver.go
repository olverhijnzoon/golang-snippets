package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Multi Channel Receiver")

	/*
		This program creates two channels (channel1 and channel2) and two goroutines are spawned. One goroutine sends random integers over channel1 every 500 milliseconds, while the other sends random integers over channel2 every 750 milliseconds. The main goroutine then enters an infinite loop and uses a select statement to receive values from either channel1 or channel2. If a value is received from channel1, it prints "Received from channel1: [value]". If a value is received from channel2, it prints "Received from channel2: [value]".
	*/

	// Create two channels to communicate between goroutines
	channel1 := make(chan int)
	channel2 := make(chan int)

	// Spawn a goroutine to send messages over channel1
	go func() {
		for {
			// Send a random integer over channel1
			channel1 <- rand.Int()
			// Wait a short time before sending the next message
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Spawn a goroutine to send messages over channel2
	go func() {
		for {
			// Send a random integer over channel2
			channel2 <- rand.Int()
			// Wait a short time before sending the next message
			time.Sleep(750 * time.Millisecond)
		}
	}()

	// Receive messages from both channels using the select statement
	for {
		select {
		case val := <-channel1:
			fmt.Println("Received from channel1:", val)
		case val := <-channel2:
			fmt.Println("Received from channel2:", val)
		}
	}
}
