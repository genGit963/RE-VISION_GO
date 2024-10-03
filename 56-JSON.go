package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
---> JSON <---

detail info: https://pkg.go.dev/encoding/json, https://go.dev/blog/json
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
	Page   int      `json:"page_no"`
	Fruits []string `json:"fruits"`
}

func learn_JSON() {
	fmt.Println("\n------------ learn_JSON --------------")

	//  encoding basic data types to JSON strings
	bloB, _ := json.Marshal(true)
	fmt.Println("Encode json: ", bloB)
	fmt.Println("Decode json: ", string(bloB), "\n")

	intB, _ := json.Marshal(1)
	fmt.Println("Encode json: ", intB)
	fmt.Println("Decode json: ", string(intB), "\n")

	fltB, _ := json.Marshal(2.34)
	fmt.Println("Encode json: ", fltB)
	fmt.Println("Decode json: ", string(fltB), "\n")

	strB, _ := json.Marshal("gopher")
	fmt.Println("Encode json: ", strB)
	fmt.Println("Decode json: ", string(strB), "\n")

	// slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("Encode json: ", slcB)
	fmt.Println("Decode json: ", string(slcB), "\n")

	mapD := map[string]string{"apple": "red", "orange": "organe", "banana": "yellow"}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("Encode json: ", mapB)
	fmt.Println("Decode json: ", string(mapB), "\n")

	/*
		The JSON package can automatically
		encode your custom data types.

		It will only include exported fields
		in the encoded output
		and will by default use those names as the JSON keys.
	*/
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("Encode json: ", res1B)
	fmt.Println("Decode json: ", string(res1B), "\n")

	/*
		You can use tags on struct field declarations
		to customize the encoded JSON key names.

		Check the definition of response2
		above to see an example of such tags.
	*/
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("Encode json: ", res2B)
	fmt.Println("Decode json: ", string(res2B), "\n")

	/*
		decoding JSON data into Go values.
	*/
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	/*
	   We can also decode JSON into custom data types.
	   This has the advantages of adding additional type-safety
	   to our programs and eliminating
	   the need for type assertions when accessing the decoded data.
	*/
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}

// func main() {
// 	learn_JSON()
// 	println("\n-------------------------------")
// }
