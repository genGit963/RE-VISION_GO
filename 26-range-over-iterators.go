package main

import (
	"fmt"
	"iter"
	"slices"
)

type Listing[T any] struct {
	head, tail *Element[T]
}

type Element[T any] struct {
	next *Element[T]
	val  T
}

func (list *Listing[T]) Push(v T) {
	if list.tail == nil {
		list.head = &Element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &Element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *Listing[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := list.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func learn_range_Over_Iterator() {
	list := Listing[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)

	// Since List.All returns an iterator, we can use it in a regular range loop.
	for e := range list.All() {
		fmt.Println(e)
	}

	// Packages like slices have a number of useful functions to work with iterators. For example, Collect takes any iterator and collects all its values into a slice.
	all := slices.Collect(list.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		// Once the loop hits break or an early return, the yield function passed to the iterator will return false.

		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

}

func main() {
	learn_range_Over_Iterator()
	println()
}
