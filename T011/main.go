package main

import "fmt"

// constants (typed and untyped)
// also for multiple constants can be declared with ()

const a = 10
const b = 10.20
const c = "Test"

const (
	x int     = 12
	y float64 = 12.1
	z string  = "Test01"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println("")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
