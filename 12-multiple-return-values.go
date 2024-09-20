package main

import "fmt"

func vals() (int, int) {
	return 4, 5
}

func learnMultipleReturnValues() int {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	return a + b
}

func main() {
	res := learnMultipleReturnValues()
	fmt.Println("Result: ", res)
}
