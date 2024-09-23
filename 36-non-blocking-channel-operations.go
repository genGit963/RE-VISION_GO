package main

import "fmt"

/*
Basic sends and receives on channels are blocking.
 However, we can use select with a default clause to
 implement non-blocking sends, receives, and even non-blocking
 multi-way selects.

*/

func learn_non_blocking_channel_operations() {
	fmt.Println("\n------------- Learn Non blocking channel operations ------------")

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received ")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message: ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message: ", msg)
	case sig := <-signals:
		fmt.Println("no activity", sig)
	default:
		fmt.Println("no activity")
	}
}

func main() {
	learn_non_blocking_channel_operations()
	println("-------------------------------")
}
