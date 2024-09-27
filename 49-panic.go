package main

import (
	"fmt"
	"os"
)

/*
A panic typically means "something went unexpectedly wrong".

Mostly we use it to "fail fast" on errors
 1. that shouldn’t occur during normal operation,
 2. or that we aren’t prepared to handle gracefully.
*/
func learn_Panic() {
	fmt.Println("\n----------- learn_Panic -----------")

	// panic("a problem") // try uncommenting and commenting

	/*
		A common use of panic is to abort if a function
		returns an error value that" we don’t know
		how to (or want to) handle."

		Here’s an example of panicking if we get an
		unexpected error when creating a new file.
	*/
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}

// func main() {
// 	learn_Panic()
// 	println("\n-------------------------------")
// }

/*
Note--->
 that unlike some languages which use exceptions
 for handling of many errors,
 in Go it is idiomatic to use error-indicating
 return values wherever possible.
*/
