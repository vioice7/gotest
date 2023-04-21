package main

import "fmt"

func main() {

	// make type lenght capacity
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 12
	x[9] = 120

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 125)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// it doubles the capacity if it runs over lenght after appending

	x = append(x, 125)

	x = append(x, 125)

	x = append(x, 125)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
