// Package foo defines an interface for a Foo type.
package foo

// Foo is an interface for a type that can perform a foo operation.
type Foo interface {
	DoFoo() error
}

/*
	"Declare an interface where it is used, not where it is implemented. Unless the interface is well discovered." https://twitter.com/inancgumus/status/1605116543553019905?s=20&t=fTo98MYIjsZ5nw44XnTmRw

	In Go, you can declare an interface in a separate package from where it is implemented. This allows you to define a set of behaviors that a type must implement, without specifying how those behaviors are implemented.
*/
