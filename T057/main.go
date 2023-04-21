package main

import (
	"fmt"
	"sort"
)

// sort package
// system int float string
// test

func main() {

	fmt.Println()

	xi := []int{3, 4, 6, 1, 2, 7, 9, 8, 5, 0}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	xf := []float64{3.12, 5.12, 6.21, 0.12, 0.34, 0.53, 1.20}

	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println(xf)
	fmt.Println()

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

	sort.Float64s(xf)
	fmt.Println(xf)

}
