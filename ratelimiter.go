package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Rate Limiter")

	/*
		Rate limiting is a technique used to control the rate at which a particular resource is accessed or used. It is often used to prevent overuse of a limited resource, such as a network connection or a server.

		In Go, rate limiting can be implemented using goroutines, channels, and tickers. A ticker is a mechanism for executing code at regular intervals, and it can be used to create a rate limiter that allows a certain number of requests per second, for example.

		To implement rate limiting in Go, you can create a ticker that sends the current time on a channel at a specified interval. You can then use this ticker to control the rate at which requests are processed by waiting for the ticker to send the current time before processing each request.

		For example, the following code creates a rate limiter that allows up to 2 requests per second.
	*/

	// Create a rate limiter that allows up to 2 requests per second.
	rateLimiter := time.NewTicker(time.Second / 2)

	// Create a channel to receive the rate-limited requests.
	requests := make(chan string)

	// Goroutine that simulates incoming requests.
	go func() {
		for {
			requests <- "Just another request!"
		}
	}()

	// Goroutine that processes the rate-limited requests.
	go func() {
		for {
			// Wait for a request to be available.
			<-requests

			// Wait for the rate limiter to allow the request.
			<-rateLimiter.C

			/*
				In the last line, the <- operator in Go is the receive operator. It is used to receive a value from a channel and to store the value in a variable. Further, the ".C" is a channel of type "<-chan time.Time" and refers to a field of the rateLimiter variable, which is of type *time.Ticker. The *time.Ticker type is a pointer to a time.Ticker struct, which represents a ticker that sends the current time on a channel at regular intervals.
			*/

			// Process the request.
			fmt.Println("Processing just another request at", time.Now())
		}
	}()

	// Sleep for a while to allow some requests to be processed.
	time.Sleep(time.Second * 5)
}
