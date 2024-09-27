package main

import (
	"fmt"
	"time"
)

/*
Go’s select lets you wait on multiple channel operations.

Combining goroutines and channels with select is
 a powerful feature of Go.

*/

func learn_select() {
	fmt.Println("\n---------- Learn Select --------- ")

	// channels
	c1 := make(chan string)
	c2 := make(chan string)

	// go routines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("iteration: ", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received : ", msg1)
		case msg2 := <-c2:
			fmt.Println("received : ", msg2)
		}
	}
}

func main() {
	learn_select()
	println()
}
