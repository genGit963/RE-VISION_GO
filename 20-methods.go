package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

/*
Methods can be defined for either pointer or value receiver types.
Here’s an example of a value receiver.
*/
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}
func studyMethods() {
	fmt.Println("---------- Learn about methods ---------")
}

func main() {
	studyMethods()
}
