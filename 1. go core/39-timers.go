package main

import (
	"fmt"
	"time"
)

func learn_Timers() {
	fmt.Println("\n-------------- Learn Timers ------------")
	/*
	 Timers represent a single event in the future.
	 You tell the timer how long you want to wait,
	 and it provides a channel that will be notified
	 at that time.
	*/
	timer1 := time.NewTimer(2 * time.Second) //  This timer will wait 2 seconds.

	/*
		The <-timer1.C blocks on the timer’s channel C until
		 it sends a value indicating that the timer fired.
	*/
	<-timer1.C

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	/*
		Give the timer2 enough time to fire,
		if it ever was going to, to show it is in fact stopped.
	*/
	time.Sleep(2 * time.Second)
}

// func main() {
// 	learn_Timers()
// 	println("\n-------------------------------")
// }
