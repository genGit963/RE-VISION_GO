package main

import (
	"fmt"
	"regexp"
)

/*
---> Regular Expression <---

Go offers built-in support for regular expressions.
Here are some examples of common regexp-related tasks in Go.

*/

func learn_regular_expression() {
	fmt.Println("\n------------ learn_regular_expression --------------")

	match, _ := regexp.MatchString("p([a-z]+)ch", "prekdjafjfajfad ach")
	fmt.Print(match)

	// r, _ := regexp
}

func main() {
	learn_regular_expression()
	println("\n-------------------------------")
}
