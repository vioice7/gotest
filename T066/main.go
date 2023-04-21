package main

import "fmt"

func main() {

	// make a channel
	c := make(chan int)

	// channels should be in a go routine (channels block)

	go func() {
		// put a value on a channel
		c <- 42
	}()

	// take value off a channel
	fmt.Println(<-c)

}
