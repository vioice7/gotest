package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())

	counter := 0

	// number of go routines
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		// anonymus function in a go routine go func(){}()
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()

		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
