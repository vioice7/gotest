package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	x := sum(xi...)

	// the operator ... after a slice transforms the slice of int into individual ints
	// we can pass 0 or more parameter
	fmt.Println("The sum of the numbers is:", x)

}

// variatic parameters are stored in a slice
// ... is in front of the type, the variatic parameter must be last

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x)) // lenght
	fmt.Println(cap(x)) // capacity

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item at index", i, "we add", v, "total is", sum)
	}
	//fmt.Println("Total is", sum)
	return sum
}
