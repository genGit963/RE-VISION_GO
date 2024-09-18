package main

import "fmt"

func if_else_study() {
	if 7%2 == 0 {
		fmt.Println("7 is even")

	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if false || true {
		fmt.Println("false || true")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}

// func main() {
// 	if_else_study()
// }
