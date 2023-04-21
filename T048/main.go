package main

import "fmt"

// callback functions example

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := sum(ii...)
	fmt.Println("all numbers sum is a:", s)

	s10 := even(sum, ii...)
	fmt.Println("even numberes sum is a:", s10)

	s11 := odd(sum, ii...)
	fmt.Println("odd numbers sum is a:", s11)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, mi ...int) int {
	var ei []int
	for _, v := range mi {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}
	return f(ei...)
}

func odd(f func(xi ...int) int, mi ...int) int {
	var ei []int
	for _, v := range mi {
		if v%2 != 0 {
			ei = append(ei, v)
		}
	}
	return f(ei...)
}
