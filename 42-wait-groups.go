package main

import (
	"fmt"
	"sync"
	"time"
)

// to wait multiple go_routines to finish, "waitgroup" is applied

// workerf2 : workerf2 function no. 2
func workerf2(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second) // indicate smth complicated caculation is done
	fmt.Printf("Worker %d done \n", id)
}

func learn_Wait_Groups() {
	fmt.Println("\n------------- learn Wait Groups -------------")

	/*
		This WaitGroup is used to wait for all the goroutines launched
		 here to finish.

		 --------- NOTE ---------
		  if a WaitGroup is explicitly passed into functions,
		  it should be DONE BY POINTER.
		 ---------------------------
	*/
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		/*
			Wrap the worker call in a closure that makes sure to tell
			the WaitGroup that this worker is done.
			This way the worker itself does not have to be
			aware of the concurrency primitives involved in its execution.
		*/

		go func() {
			defer wg.Done()
			workerf2(i)
		}()
	}
	/*
		Block until the WaitGroup counter goes back to 0;
		 all the workers notified they’re done.
	*/
	wg.Wait()
}

func main() {
	learn_Wait_Groups()
	println("\n-------------------------------")
}
