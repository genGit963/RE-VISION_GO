package main

import "fmt"

func SliceIndex[Slice ~[]Element, Element comparable](s Slice, v Element) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func learn_Generics() {
	fmt.Println("\n-------- Learn Generics ----------")
}

func main() {
	learn_Generics()
	println()
}
