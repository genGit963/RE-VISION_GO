package main

import "fmt"

func forLoopStudy() {

	// method 1
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// method 2
	for j := 0; j < 3; j++ {
		fmt.Println("range: ", j)
	}

	// method 3
	for i := range 3 {
		fmt.Println("range: ", i)
	}

	// method 4
	for {
		fmt.Println("loop")
		break
	}

	// method 4
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("odd:", n)
	}

}

// func main() {
// 	forLoopStudy()
// }
