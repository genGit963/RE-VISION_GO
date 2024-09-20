package main

import "fmt"

/*
We’ll show how pointers work in contrast to values with 2 functions:
zeroval and zeroptr.
zeroval has an int parameter, so arguments will be passed to it by value.
zeroval will get a copy of ival distinct from the one in the calling function.
*/
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println("see iptr parameter \n value: ", iptr, "\n self adrs:", &iptr)
	*iptr = 0
}

func learn_Pointers() {
	fmt.Println("\n---------- Learning about pointers ----------")

	var num int = 42                                                 // A normal integer variable
	fmt.Println("Before:", num)                                      // Output: Before: 42
	zeroptr(&num)                                                    // Pass the address of num to the function (Referencing)
	fmt.Println("After passing: ", &num, "then value of num: ", num) // Output: After: 0
}

func main() {

	learn_Pointers()

	fmt.Println("Summary \n Referencing (&): Gets the address of a variable (e.g., &num). \nDereferencing (*): Accesses or modifies the value at that address (e.g., *iptr).")
}
