package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Snippets!")
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## HTTP Server")

	/*
		This is a simple example of a HTTP server in Go.
	*/

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
