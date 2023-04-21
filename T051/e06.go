package main

import "fmt"

func main() {

	// anonymous function func(){}() and can be used with go routines

	func() {
		fmt.Println("An anonymous function has ran!")
	}()

}
