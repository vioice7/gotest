package main

import "fmt"

func main() {

	// func expressions (assign a function to a variable)
	f00 := func() {
		fmt.Println("my func expression variable")
	}
	// run that variable as a function
	f00()

	f11 := func(x int) {
		fmt.Println("test", x)
	}

	f11(12)
}
