package main

import (
	"fmt"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func handler(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex to prevent concurrent access to the counter variable
	mutex.Lock()
	counter++
	mutex.Unlock()
	fmt.Fprintf(w, "You are visitor number %d", counter)
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## HTTP Server")

	/*
		This is an example of a HTTP server in Go.
	*/

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
