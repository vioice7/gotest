package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		// do a for looping and select all the values from c and then q
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			// fmt.Println(v)
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		// we need to send a value to q as well
		q <- 1
		close(c)
	}()

	return c
}
