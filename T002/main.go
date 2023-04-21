package main

import "fmt"

// the scope of a variable (wher it is visible)

// var can be used to declare
// a variable outside a funcion
var w = 43

// declare a variable z with type int (default is 0)
var z int

func main() {

	// short declaration operator :=
	// tha variable is already declared whit this operator
	// assign a value of a certain type

	// short declaration can only be used to declare
	// a variable inside a function
	x := 12

	fmt.Println("x =", x)

	x = 1200

	fmt.Println("x =", x)

	// statement
	y := 12 + x

	fmt.Println("y =", y)

	fmt.Println("z =", w)

}
