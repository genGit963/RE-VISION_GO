package main

import "fmt"

func learn_Channel_buffering() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

func main() {
	learn_Channel_buffering()
	println()
}
