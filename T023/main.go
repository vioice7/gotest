package main

import "fmt"

func main() {

	// array has to have a specified size
	// the lenght is part of it's type
	var x [5]int

	fmt.Println(x)

	x[4] = 12

	fmt.Println(x)

	fmt.Println(len(x))

}
