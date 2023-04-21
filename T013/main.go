package main

import "fmt"

var x = 2

func main() {

	fmt.Printf("%d\t\t%b", x, x)

	fmt.Printf("\n")

	x = 4
	fmt.Printf("%d\t\t%b", x, x)

	fmt.Printf("\n")

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)

}
