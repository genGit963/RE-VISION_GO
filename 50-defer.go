package main

import (
	"fmt"
	"os"
)

/*

Defer is used to ensure that a function call is
"performed later in a program’s execution",
usually for purposes of cleanup.

defer is often used where
e.g. ensure and finally would be used in other languages.

*/

func learn_defer() {
	fmt.Println("\n-------- learn_defer --------")

	/*
		Suppose we wanted to create a file,
		write to it, and then close when we’re done.

		Here’s how we could do that with defer.
	*/
	f := createFile("defer.txt")

	/*
		Immediately after getting a file object with createFile,
		we defer the closing of that file with closeFile.

		This will be executed at the end of the enclosing function (learn_defer),
		after writeFile has finished.
	*/

	defer closeFile(f)
	writeFile(f)

}

func createFile(s string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "File is written now with data. !")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %V\n", err)
		os.Exit(1)
	}
}

// func main() {
// 	learn_defer()
// 	println("\n-------------------------------")
// }
