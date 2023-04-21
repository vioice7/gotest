package main

import "fmt"

func main() {

	// channels allow one value to sit in there (this works as well)

	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
