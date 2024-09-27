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

	/*
		We can use the same technique to sort a slice of values
		that aren’t built-in types.
	*/
	type Person struct {
		name   string
		height float32
		age    int
	}

	people := []Person{
		Person{name: "Amen", height: 4.5, age: 14},
		Person{name: "Ern", height: 6.5, age: 24},
		Person{name: "Nin", height: 2.5, age: 4},
		Person{name: "Eopn", height: 5.53, age: 44},
		Person{name: "Deen", height: 4.5, age: 22},
		Person{name: "Gueen", height: 5.05, age: 54},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("sort by Age: ", people)

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.height, b.height)
	})
	fmt.Println("sort by height: ", people)

}

func main() {
	learn_sorting_by_functions()
	println("\n-------------------------------")
}
