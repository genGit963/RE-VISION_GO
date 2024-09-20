package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, ", ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("Total Sum: ", total)
	return total
}

func learnVariadicFunc() int {
	return sum(1, 2, 4, 5)
}

func main() {
	res := learnVariadicFunc()
	fmt.Println("Finally, ", res)
}
