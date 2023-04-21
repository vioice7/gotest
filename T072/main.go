package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		c <- 10
		close(c)
	}()
	// the comma ok idiom
	v, ok := <-c
	fmt.Println(v, ok)
	// pull  from channel after closing
	v, ok = <-c
	fmt.Println(v, ok)
}
