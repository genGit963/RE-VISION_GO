package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working......")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func learn_Channel_Synchronization() {
	done := make(chan bool, 1)
	go worker(done)

	if <-done {
		fmt.Println("All work DONE !!")
	}
}

func main() {
	learn_Channel_Synchronization()
	println()
}
