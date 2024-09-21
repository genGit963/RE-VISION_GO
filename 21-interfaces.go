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

}

func main() {
	learn_interfaces()
	println()
}
