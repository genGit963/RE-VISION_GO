package main

import (
	"fmt"
	"time"
)

func oneFunction(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func learn_GoRoutine() {
	// Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
	oneFunction("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go oneFunction("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")
}

func main() {
	learn_GoRoutine()
	println()
}
