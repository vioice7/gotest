package main

import "fmt"

type vartipulmeu int

var x vartipulmeu

func main() {

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)

}
