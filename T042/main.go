package main

import "fmt"

// defer the execution of a function

func main() {
	// defer runs when the containing function exits

	defer test0()
	test1()

}

func test0() {

	fmt.Println("test0")

}

func test1() {

	fmt.Println("test1")

}
