// array is collection of fixed number of items. However, slices are much more common

package main

import "fmt"

func learnArray() {

	var a [5]int
	fmt.Println("study array a: ", a)

	a[4] = 10
	fmt.Println("assign a[4]: ", 10, "Access a[4]: ", a[4])

	fmt.Println("len of a: ", len(a))

	b := [5]int{2, 3, 4, 5, 6} // declare and initialize the array
	fmt.Println("checkout array b: ", b)

	b = [...]int{1, 2, 3, 4, 5} // you can also have the compiler count the number of elements for you with ...
	fmt.Println("dcl: ", b)

	b = [...]int{1, 3: 10, 3} // If you specify the index with :, the elements in between will be zeroed.
	fmt.Println("idx certain of b:", b)

	var twoD [2][3]int
	var i = 2
	var j = 3
	// initialization of twoD array
	for i < 2 {
		for j < 3 {
			twoD[i][j] = i * j
		}
	}
	fmt.Println("2d: ", twoD)

	// var threeD = [2][3][4]int //3d array may not supported
	// fmt.Println("2d: ", threeD)

	// multidimensional array declare and initialize
	twoD = [2][3]int{
		{1, 4, 5}, {1, 3, 9},
	}
	fmt.Println("created again twoD and initialized: ", twoD)

}

// func main() {
// 	learnArray()
// }
