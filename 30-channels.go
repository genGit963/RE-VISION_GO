package main

import "fmt"

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and

	receive those values into another goroutine.
*/
func spacing() {
	fmt.Println("\n---------- Learn Channels ----------")
	fmt.Println()
}

func learn_Channels() {
	spacing()

	/*
		Create a new channel with make(chan val-type).
		Channels are typed by the values they convey.
	*/
	messages := make(chan string)

	go func() {
		messages <- "ping go"
	}()

	msg := <-messages
	fmt.Println(msg)

}

// func main() {
// 	learn_Channels()
// 	println()
// }
