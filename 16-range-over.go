package main

import "fmt"

func learn_Range_Over() {

	//  range can be use to iterate through list
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum : ", sum)

	// also provide index of array or slices
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range for maps, destructure every key and value
	kvs := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range kvs {
		fmt.Printf("key: %s -> value:%s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("keys: ", k)
	}

	// range on strings iterates over Unicode code points.
	for i, c := range "go Cool 😜" {
		fmt.Printf("index i has: %d character: %c\n", i, c)
	}

}

func main() {
	learn_Range_Over()
}
