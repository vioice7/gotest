package main

import "fmt"

func main() {

	x := []int{0, 1, 2, 3, 4}

	fmt.Println(x)
	fmt.Println(x[4])
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// index, value format
	for i, v := range x {
		fmt.Println(i, v)
	}

}
