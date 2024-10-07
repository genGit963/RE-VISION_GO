package main

import (
	"fmt"
	"os"
)

/*
---> Exit <---
Use os.Exit to immediately exit with a given status.
*/

func learn_exit() {
	fmt.Println("\n------------ learn_exit --------------")
	/*
		defers will not be run when using os.Exit,
		so this fmt.Println will never be called.
	*/
	defer fmt.Println("!")
	// Exit with status 3.

	os.Exit(3)

	/*
		Note :
			that unlike e.g. C, Go does not use an
			integer return value from main to indicate exit status.

			If you’d like to exit with a
			non-zero status you should use os.Exit.
	*/
}

func main() {
	learn_exit()
	println("\n-------------------------------")
}
