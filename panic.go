package main

import (
	"fmt"
	"io"
	"os"
)

func division(x, y int) int {
	if y == 0 {
		// Panic with a error message.
		panic("division by zero")
	}
	return x / y
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		// If there is an error, panic.
		panic(err)
	}

	// Register a deferred function to close the file.
	defer func() {
		if err := file.Close(); err != nil {
			// If there is an error closing the file, panic.
			panic(err)
		}
	}()
	return file
}

func main() {
	fmt.Println("### Golang Snippets ### Panic")
	/*
		A panicking function is a Go function that has encountered an exceptional condition and is unwinding the call stack to find a recoverable point.

		When a function panics, it stops executing and starts propagating the panic to its caller. If the caller is not prepared to handle the panic, it will also panic, and so on, until the panic reaches the top-level of the program or a recoverable point is found.

		The recover() function can be used to recover from a panic. It should be called from a deferred function, and it will return the value that was passed to panic(). If recover() is called outside the context of a panicking function, it will return nil.

		It is recommended to use recover() only as a last resort, to prevent a program from crashing. It is better to handle errors using Go's built-in error handling mechanisms, such as the error type and the try-catch-like if-err idiom.
	*/

	// If panic, the deferred function will retrieve the panic value using the recover() function.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Second call will panic - active to test first example.
	fmt.Println(division(10, 2))
	//fmt.Println(division(10, 0))

	/*
		When a function panics, any deferred functions that have been registered by the panicking
		function will still be executed before the panic is propagated to the caller. This can be
		useful for cleaning up resources, logging information, or performing other tasks that need to be done before the panic is handled.

		In this example, the deferred function in openFile() is used to clean up the file resource
		before the panic is propagated. If the openFile() function panics, the deferred function will be called and will close the file. This ensures that the file is always closed, even if the openFile() function panics. Similarly, if the main() function panics, the deferred function in main() will be called and will handle the panic, allowing the program to continue executing.
	*/

	// Call openFile() and recover from any panic that occurs.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	file := openFile("myfile.txt")
	_, err := io.Copy(os.Stdout, file)
	if err != nil {
		// If there is an error, panic.
		panic(err)
	}
}
