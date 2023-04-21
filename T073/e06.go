package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			// to push  values to the channel
			c <- i
		}
		// close the channel
		close(c)

	}()

	for i := 0; i < 100; i++ {
		// to pull the values from the channel
		fmt.Println(<-c)
	}
}
