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

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range for maps
	kvs := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range kvs {
		fmt.Printf("key: %s -> value:%s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys: ", k)
	}

	for i, c := range "go lang is great" {
		fmt.Printf("index i has: %d character: %c", i, c)
	}

}

func main() {
	learn_Range_Over()
}
