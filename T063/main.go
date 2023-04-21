package main

import (
	"fmt"
	"runtime"
	"sync"
)

// mutex lock used to lock the acces of other
// goroutines to a certain variable

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())

	counter := 0

	// number of go routines
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		// anonymus function in a go routine go func(){}()
		go func() {
			// code is locked so nobody except the counter
			// can access the variable
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
