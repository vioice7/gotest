package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// using atomic to syncronise go routines

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())

	var counter int64

	// number of go routines
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		// anonymus function in a go routine go func(){}()
		go func() {
			// code is locked so nobody except the counter
			// can access the variable
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
