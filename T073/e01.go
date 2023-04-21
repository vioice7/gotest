package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 12
	}()

	fmt.Println(<-c)

	// for buffered channel you have to add to the stack
	// how many channels to be added into the buffer
	// c := make(chan int, 1)

	// add a value to that buffer
	// c <- 12

	// take the value from that buffer
	// fmt.Println(<-c)

}
