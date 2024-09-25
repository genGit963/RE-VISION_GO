package main

import (
	"fmt"
	"time"
)

func oneWorker(done chan bool) {
	fmt.Println("Working......")
	time.Sleep(time.Second)
	fmt.Println("done", done)

	done <- true
	fmt.Println("end", done)
}

func learn_Channel_Synchronization() {
	done := make(chan bool, 1)
	go oneWorker(done)

	if <-done {
		fmt.Println("All work DONE !!")
	}
}

// func main() {
// 	learn_Channel_Synchronization()
// 	println()
// }
