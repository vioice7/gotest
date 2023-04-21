package main

import "fmt"

func main() {

	x := 42

	y := "James Bond"

	z := true

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v %v %v", x, y, z)
	fmt.Println()
	fmt.Printf("%T %T %T", x, y, z)

}
