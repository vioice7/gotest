package main

import "fmt"

type vartipulmeu int

var x vartipulmeu

var y int

func main() {

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)

	y = int(x)

	fmt.Printf("%v\n", y)
	fmt.Printf("%T", y)

}
