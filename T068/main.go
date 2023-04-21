package main

import "fmt"

func main() {

	// send and receive channel
	c := make(chan int, 3)

	// channel send only
	cs := make(chan<- int, 3)

	// channel receive only
	cr := make(<-chan int, 3)

	fmt.Printf("channel directional: %T\n", c)
	fmt.Printf("channel send only: %T\n", cs)
	fmt.Printf("channel receive only: %T\n", cr)

}
