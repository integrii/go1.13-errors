package main

// Playing around with errors in the go 1.13 release today.
// Helpful blog: https://crawshaw.io/blog/xerrors
// Godoc for xerrors package (godoc.org not updated yet): https://godoc.org/golang.org/x/xerrors#example-New--Errorf
// Go 1.13 changelog: https://tip.golang.org/doc/go1.13

import (
	"errors"
	"fmt"
)

func main() {

	// do something that causes an error
	err := doBadCodepath()
	fmt.Println("Plain println:")
	fmt.Println(err)

	// check if this is an errorSomeError
	if errors.Is(err, errorSomeError) {
		fmt.Println("Error is of type SomeError")
	} else {
		fmt.Println("Error is NOT a SomeError")
	}

	// unwrap the error and output the one inside
	fmt.Println("Unwrapped:")
	fmt.Println(errors.Unwrap(err))

}

// doBadCodepath tries to do something that returns a defined
// error and hands that error back to us
func doBadCodepath() error {
	err := causeError()
	return fmt.Errorf("a bad codepath was called: %w", err)
}

// causeError returns a specific error of a known type
func causeError() error {
	return errorSomeError
}

// SomeError defines a specific error that may have various
// pieces of deocration (metadata) attached to it
type SomeError struct {
	text string
}

// Error implements the error interface on SomeError
func (s SomeError) Error() string {
	return s.text
}

// errorSomeError represents the exact error that will be generated downstream
// and then checked against later.
var errorSomeError = SomeError{text: "This is some error that happened down the stack.  A specific one at that."}
