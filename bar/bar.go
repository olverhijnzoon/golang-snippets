// Package bar provides an implementation of the Foo interface
package bar

import (
	"fmt"

	"github.com/olverhijnzoon/golang-snippets/foo"
)

// MyFoo is a type that implements the Foo interface
type MyFoo struct{}

// DoFoo performs the foo operation.
func (mf *MyFoo) DoFoo() error {
	// Implementation goes here.
	fmt.Printf("Yicks!")
	return nil
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Foo Interface Bar")
	foo.DoFoo()
}
