package main

import "fmt"

// bool value check
var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 17
	b := 42
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
}
