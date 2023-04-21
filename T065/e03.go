package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// create a var for wait group
	var wg sync.WaitGroup

	// create a variable for incrementing
	incrementor := 0
	// create a variable for number of go routines
	gs := 100
	// add go routines to wait group
	wg.Add(gs)

	// create a loop for all the go routines
	for i := 0; i < gs; i++ {
		// create a go routine
		// with an anonymous function
		go func() {
			// set a variable to the incrementor value
			v := incrementor
			// run gosched to yeald the processor
			runtime.Gosched()
			// increment the variable
			v++
			// assign to the incrementor the value v
			incrementor = v

			//check the value of the incrementor
			fmt.Println(incrementor)

			// notify the wait group that
			// each go routine is done
			wg.Done()
		}()
	}
	// wait for all go routines to be done
	wg.Wait()
	fmt.Println("The value:", incrementor)
}
