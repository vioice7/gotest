package main

import (
	"fmt"
	"runtime"
)

// variable types

var x int
var y float64

func main() {
	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println("")

	c0 := 10
	c1 := 10.20

	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Printf("%T\n", c0)
	fmt.Printf("%T\n", c1)

	fmt.Println("")

	fmt.Println("Arhitectura sistemului de operare:")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
