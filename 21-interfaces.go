package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

/*
To implement an interface in Go,
we just need to implement all the methods in the interface.
Here we implement geometry on rects
*/
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

/*
implementation for circles
*/
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

// measure
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func learn_interfaces() {
	fmt.Println("\n----------- Learn about interfaces ----------")
	rt := rectangle{width: 3, height: 5}
	cle := circle{radius: 7}

	measure(rt)
	measure(cle)
	fmt.Println("reactange: ", rt)
	fmt.Println("circle: ", cle)

	// try directly the interface
	// var gmtryRect geometry = rt
	var gmtryRect geometry = rectangle{width: 40, height: 20}
	fmt.Println("reactange geometry area: ", gmtryRect.area())
	fmt.Println("reactange geometry perimeter: ", gmtryRect.perimeter())

	// var gmtryCircle geometry = cle
	var gmtryCircle geometry = circle{radius: 40}
	fmt.Println("circle geometry area: ", gmtryCircle.area())
	fmt.Println("circle geometry perimeter: ", gmtryCircle.perimeter())

	// conclusion
	fmt.Println("Key Takeaways:\n" +
		"• Structs are used to store concrete data with predefined fields.\n" +
		"• Interfaces define abstract behaviors that types can implement, enabling polymorphism.\n" +
		"• Structs and interfaces can work together, where a struct can implement one or more interfaces by defining the required methods.\n" +
		"• A key feature of interfaces is that they allow different types (like Dog, Car, or Person) to implement the same behavior and be used interchangeably as the same interface type.")

}

func main() {
	learn_interfaces()
	println()
}
