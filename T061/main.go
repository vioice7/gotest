package main

import (
	"fmt"
	"runtime"
	"sync"
)

// variable for wait group
// syncronisation primitive built in a code
var wg sync.WaitGroup

func main() {

	fmt.Println("Sistem de operare:\t", runtime.GOOS)
	fmt.Println("Arhitectura:\t\t", runtime.GOARCH)
	fmt.Println("Procesoare:\t\t", runtime.NumCPU())
	fmt.Println("Rutine Go:\t\t", runtime.NumGoroutine())

	// wait group wait for adding 1 thing
	// add 1 thing to wait for
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("Rutine Go:\t\t", runtime.NumGoroutine())
	// wait until all things added are done
	// when all waits are done our program can exit
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// tells the wait group variable that this function is done
	// remove 1 thing to wait for
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
