package main

import "fmt"

// function that returns a type func of type int
// assign it to a variable
// shows the number

func test0() func() int {
	return func() int {
		return 12
	}
}

func main() {
	f := test0()
	fmt.Println("adress:", f)
	fmt.Println(f())
}
