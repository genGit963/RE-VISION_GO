package main

import (
	"crypto/sha256"
	"fmt"
)

/*
---> SHA256 Hashes<---

SHA256 : https://en.wikipedia.org/wiki/SHA-2

hashes are frequently used to compute short identities
for binary or text blobs.

using built in library "crypto/sha256"

For example,
TLS/SSL certificates use SHA256 to compute a certificate’s signature.
Here’s how to compute SHA256 hashes in Go.
*/

func learn_sha256_hashes() {
	fmt.Println("\n------------ learn_sha256_hashes --------------")

	s := "sha256 will do ciphering the text"
	h := sha256.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	h.Write([]byte(s))

	/*
		This gets the finalized hash result as a byte slice.
		The argument to Sum can be used to append to
		an existing byte slice: it usually isn’t needed.
	*/
	bs := h.Sum(nil)

	fmt.Println("plain: ", s)
	fmt.Printf("hash: %x\n", bs)
}

func main() {
	learn_sha256_hashes()
	println("\n-------------------------------")
}
