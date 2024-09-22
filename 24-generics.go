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

/*
We can define methods on generic types just like we do on regular types,
but we have to keep the type parameters in place. The type is List[T],
not List.
*/
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

/*
AllElements returns all the List elements as a slice.
In the next example we’ll see a more idiomatic way of iterating
over all elements of custom types.
*/
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func learn_Generics() {
	fmt.Println("\n-------- Learn Generics ----------")
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SliceIndex(s, "foo"))

	_ = SliceIndex[[]string, string](s, "zoo")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}

func main() {
	learn_Generics()
	println()
}
