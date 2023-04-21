package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println("The sum of the numbers is:", x)
}

// variatic parameters are stored in a slice
// ...

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item at index", i, "we add", v, "total is", sum)
	}
	//fmt.Println("Total is", sum)
	return sum
}
