package main

import (
	"fmt"
)

func main() {
	// function test1 takes parameter function test0
	test1(test0)
}

func test0() int {
	return 1000
}

// the argument of function test1 is f of type func()
// that has return type int
func test1(f func() int) {
	// prints the value from the argument
	fmt.Println(f())
}
