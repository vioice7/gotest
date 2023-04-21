package main

import "fmt"

func main() {

	// appending to a slice

	x := []int{11, 23, 45, 67, 89, 90}
	fmt.Println(x)
	x = append(x, 12, 14, 16, 18, 20)
	fmt.Println(x)

	y := []int{100, 200, 300}

	// using ... after you append all elements in y
	// the symbol ... used in front symbolise
	// a number of unlimited parameters

	x = append(x, y...)
	fmt.Println(x)
}
