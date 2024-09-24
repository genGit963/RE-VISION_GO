package main

import "fmt"

func learn_Range_Over_Channels() {

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	close(queue)
	// queue <- "five"

	/* This range iterates over each element as it’s
	received from queue. Because we closed the channel
	 above, the iteration terminates after receiving
	  the 2 elements.
	*/
	fmt.Println("range over channesls: ")
	for elem := range queue {
		fmt.Println(elem)
	}

}

func main() {
	println("\n-------------------------------")
	learn_Range_Over_Channels()
	println("\n-------------------------------")
}
