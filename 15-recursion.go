package main

import "fmt"

func factorial(n int) int {
	fmt.Print(n, " ")
	if n == 0 || n == 1 {
		return n
	} else {
		return n * factorial(n-1)
	}
}

func main() {

	factValue := factorial(10)

	fmt.Println("Factorial result: ", factValue)
}
