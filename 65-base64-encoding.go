package main

import (
	b64 "encoding/base64"
	"fmt"
)

/*
---> Base64 encoding <---
built-in support for base64 encoding/decoding
*/

func learn_base64_encoding() {
	fmt.Println("\n------------ learn_base64_encoding --------------")
	// string we’ll encode/decode.
	data := "abc$1@#$231)-=@1"
	/*
	   Go supports both standard and URL-compatible base64.

	   Here’s how to encode using the standard encoder.
	   The encoder requires a []byte so we convert our string to that type.
	*/
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("base64 encode: ", sEnc)

	/*
		Decoding may return an error, which you can check
		if you don’t already know the input to be well-formed.
	*/
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("decode:", string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}

func main() {
	learn_base64_encoding()
	println("\n-------------------------------")
}
