package main

import "fmt"

// Go supports methods defined on struct types.
type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

/*
Methods can be defined for either pointer or value receiver types.
Here's an example of a value receiver.
*/
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func studyMethods() {
	fmt.Println("\n---------- Learn about methods ---------")

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := &r
	fmt.Println("by ptr, area: ", rp.area())
	fmt.Println("by pty, perimeter: ", rp.perimeter())

	// conclusion
	fmt.Println("\nConclusion\n\nUse pointer receivers (*Type) when you: \n\tNeed to modify the struct. \n\tWant to avoid copying large structs for performance reasons.\nUse value receivers (Type) when:\n\tYou don't need to modify the struct.\n\tThe struct is small, and copying it doesn't impact performance.")
}

func main() {
	studyMethods()
	fmt.Println()
}
