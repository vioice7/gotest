package main

import "fmt"

func main() {

	// assign a func as a variable and make it ran

	a := func() {
		fmt.Println("test")
	}
	a()

	// check the type of a
	fmt.Printf("Type of a() : %T\n", a)

	b := a

	b()

	// check the type of a
	fmt.Printf("Type of b() : %T\n", b)

}
