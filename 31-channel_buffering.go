package main

import "fmt"

func learn_Channel_buffering() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	messages <- "channel msg 2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

func main() {
	go learn_Channel_buffering()
	println()
}
