package main

import "fmt"

func main() {

	a0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b0 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	a := foo(a0...)
	b := bar(b0)

	fmt.Println(a)
	fmt.Println(b)

}

// unfurl the variatic parameter to be an int slice
func foo(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}

// takes directly some int slices
func bar(x []int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}
