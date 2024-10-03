package main

import (
	"encoding/json"
	"fmt"
)

/*
---> JSON <---
Go offers built-in support for JSON encoding and decoding,
including to and from built-in and custom data types.
*/

type response1 struct {
	Page   int
	Fruits []string
}

/*
Only exported fields will be encoded/decoded in JSON.
Fields must start with capital letters to be exported.
*/
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func learn_JSON() {
	fmt.Println("\n------------ learn_JSON --------------")

	//  encoding basic data types to JSON strings
	bloB, _ := json.Marshal(true)
	fmt.Println(string(bloB))

}

func main() {
	learn_JSON()
	println("\n-------------------------------")
}
