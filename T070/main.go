package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		// put in a value into the channel
		for i := 0; i < 100; i++ {
			c <- i
		}
		// must close the channel or the reciver will still wait
		close(c)
	}()

	for v := range c {
		// pulls out a value in range
		fmt.Println(v)
	}

	fmt.Println("Abort to exit")
}
