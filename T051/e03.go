package main

import "fmt"

func main() {

	// this will run last because it executes on main func exit
	defer a0()

	// this will run first
	a1()

}

func a0() {

	fmt.Println("a0 ran")

}

func a1() {

	// the anonymous function ran at the exit of a1 func first last
	defer func() {
		fmt.Println("anonymous a01")
	}()

	fmt.Println("a1 ran")

}
