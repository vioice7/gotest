package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go t0()
	wg.Add(1)
	go t1()

	fmt.Println("Rutine Go:\t\t", runtime.NumGoroutine())
	wg.Wait()

}

func t0() {
	for i := 0; i < 100; i++ {
		fmt.Println("Test0", i)
	}
	wg.Done()
}

func t1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Test1", i)
	}
	wg.Done()
}
