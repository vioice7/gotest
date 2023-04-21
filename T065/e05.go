package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// create a var for wait group
	var wg sync.WaitGroup

	// create a variable for incrementing
	// when using atomic package
	// the variable should be in64
	var incrementor int64

	// create a variable for number of go routines
	gs := 100
	// add go routines to wait group
	wg.Add(gs)

	// create a loop for all the go routines
	for i := 0; i < gs; i++ {
		// create a go routine
		// with an anonymous function
		go func() {
			// add to the incrementor variable
			// we add directly to the memory location
			atomic.AddInt64(&incrementor, 1)

			// check the value of the incrementor
			// using atomic load
			fmt.Println("counter:", atomic.LoadInt64(&incrementor))

			// notify the wait group that
			// each go routine is done
			wg.Done()
		}()
	}
	// wait for all go routines to be done
	wg.Wait()
	fmt.Println("The value:", incrementor)
}
