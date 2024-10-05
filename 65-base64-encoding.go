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

}

func main() {
	learn_base64_encoding()
	println("\n-------------------------------")
}
