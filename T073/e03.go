package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	// use func literal to send this go routine
	go func() {

		for i := 0; i < 100; i++ {
			// add values to the channel
			c <- i
		}
		// close the channel
		close(c)

	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
