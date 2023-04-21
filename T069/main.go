package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("Abort to exit")
}

// send
func foo(c chan<- int) {
	// scope of c only for this function
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
