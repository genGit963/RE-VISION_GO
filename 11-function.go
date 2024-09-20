package main

import "fmt"

// learning function

func learnVoidFunc() {
	fmt.Println("I am learning void function")
}

func addFunc(a, b int) int {
	return a + b
}

//	var subValue = subFunc(a,b int)uint{
//		return a - b
//	}
func main() {
	// void function
	learnVoidFunc()
	// default function
	res := addFunc(10, 4)
	fmt.Println("number is : ", res)
}
