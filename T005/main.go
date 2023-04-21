package main

import "fmt"

// different types declared a different from b

var a int

// create your own type
type aim int

var b aim

func main() {

	a = 0
	b = 1

	fmt.Printf("Type of a with value %v is %T\n", a, a)
	fmt.Printf("Type of b with value %v is %T\n", b, b)

	// a = b can't be used
	// use conversion (not casting)

	a = int(b)

	fmt.Printf("Type of a with value %v is %T\n", a, a)
	fmt.Printf("Type of b with value %v is %T\n", b, b)

	b = 10
	a = int(b)

	fmt.Printf("Type of a with value %v is %T\n", a, a)
	fmt.Printf("Type of b with value %v is %T\n", b, b)

}
