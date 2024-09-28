package main

import "fmt"

/*
Go makes it possible to recover from a panic,
by using the recover built-in function.

A recover can stop a panic from aborting the program and
let it continue with execution instead.
*/

func mayPanic() {
	panic("a problem")
}

func learn_Recover() {
	fmt.Println("\n---------------- learn_Recover ------------")

	/*
		recover must be called within a deferred function.

		When the enclosing function panics, the defer will activate
		and a recover call within it will catch the panic.
	*/
	defer func() {
		if r := recover(); r != nil { //The return value of recover is the error raised in the call to panic.
			fmt.Println("Recovered Error: \n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic !")
}
func main() {
	learn_Recover()
	println("\n-------------------------------")
}
