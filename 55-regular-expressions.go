package main

import (
	"bytes"
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

	match, _ := regexp.MatchString("p([a-z]+)ch", "prekdjafjfajfadach")
	fmt.Println(match)

	/*
		Above we used a string pattern directly,
		but for other regexp tasks

		you’ll need to Compile an optimized Regexp struct.
	*/
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println("idx: ", r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}

func main() {
	learn_regular_expression()
	println("\n-------------------------------")
}
