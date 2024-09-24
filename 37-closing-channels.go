package main

import "fmt"

func learn_Closing_Channels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs
	for j := 1; j <= 60; j++ {
		jobs <- j
		fmt.Println("sent job : ", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done // await the worker using synchronizaton

	_, ok := <-jobs
	fmt.Println("received more jobs: ", ok)
}

func main() {
	println("\n-------------------------------")
	learn_Closing_Channels()
	println("\n-------------------------------")
}
