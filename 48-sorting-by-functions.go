package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
Sometimes we’ll want to sort a collection by something other
than its natural order.

For example,

	suppose we wanted to sort strings by their length
	instead of alphabetically.

Here’s an example of "custom sorts" in Go.
*/
func learn_sorting_by_functions() {
	fmt.Println("\n-------- learn_sorting_by_functions -------")

	fruits := []string{"preach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	/*
		Now we can call slices.SortFunc
		with this custom comparison function
		to sort fruits by name length.

	*/
	slices.SortFunc(fruits, lenCmp)
	fmt.Println("sorted fruits by len: ", fruits)

}

func main() {
	learn_sorting_by_functions()
	println("\n-------------------------------")
}
