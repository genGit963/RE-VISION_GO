package main

import (
	"fmt"
	"unicode/utf8"
)

/*
A Go string is a read-only SLICE of bytes.
The language and the standard library treat strings specially -

	as containers of text encoded in UTF-8.

In other languages, strings are made of “characters”.
In Go, the concept of a character is called a rune

		rune - it’s an integer that represents a Unicode code point.
	 This Go blog post is a good introduction to the topic.
*/
func examineRune(r rune) {
	/*
		Values enclosed in single quotes are rune literals.
		We can compare a rune value to a rune literal directly.
	*/
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func learn_StringAndRunes() {
	fmt.Println("\n--------- strings and runes -------")

	/*
		s is a string assigned a literal value representing the word
		“hello” in the Thai language.
		Go string literals are UTF-8 encoded text.
	*/
	// var s = "สวัสดี" // Thai
	var s = "नमस्कार" // nepali

	/*
		Since strings are equivalent to []byte,
		this will produce the length of the raw bytes stored within
	*/

	fmt.Println("Len:", len(s), s)
	s += "😜"
	fmt.Println("Len:", len(s), s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// This demonstrates passing a rune value to a function.

		examineRune(runeValue)
	}
}

func main() {
	learn_StringAndRunes()
}
