package main

import "fmt"

/*
Go’s structs are TYPED collections of FIELDS.
They’re useful for grouping data together to form records.
*/

// structs declaration
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	/*
		Go is a garbage collected language;
		you can safely return a pointer to a local variable -
		 	it will only be cleaned up by the garbage collector
			when there are no active references to it.
	*/
	p := person{name: name}
	p.age = 42
	return &p
}

func learn_structs() {
	fmt.Println("\n --------- learn Structs ---------")

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40}, (&person{name: "Ann"}).name)

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))
	fmt.Println("newPerson name: ", newPerson("Jon").name)
	fmt.Println("newPerson age: ", newPerson("Jon").age)
}

func main() {
	learn_structs()
	println()
}
