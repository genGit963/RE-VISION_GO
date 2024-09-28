package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

/*
The standard library’s strings package provides
many useful string-related functions.
Here are some examples to give you a sense of the package.
*/
func learn_String_Functions() {
	fmt.Println("\n--------------- learn String function ---------------")

	// try out all string functions
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

func main() {
	learn_String_Functions()
	println("\n-------------------------------")
}
