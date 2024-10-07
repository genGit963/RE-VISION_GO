package main

import (
	"fmt"
	"slices"
)

/*
Go’s slices package implements sorting for builtins and user-defined
types.
We’ll look at sorting for builtins first.
*/

func learn_Sorting() {
	fmt.Println("\n-------- learn_Sorting ---------")
	/*
		Sorting functions are generic, and work for any
		ordered built-in type.
		For a list of ordered types, see cmp.Ordered.
	*/
	strs := []string{"c", "b", "g", "d", "a", "e"}
	slices.Sort(strs)
	fmt.Println("Strings inorder: ", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints inorder: ", ints)

	// We can also use the slices package to check if a slice is already in sorted order.
	s := slices.IsSorted(ints)
	s1 := slices.IsSorted([]int{4, 5, 5, 2})
	fmt.Println("IsSorted: ", s, s1)
}

// func main() {
// 	learn_Sorting()
// 	println("\n-------------------------------")
// }
