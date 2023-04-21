package main

import "fmt"

// recursion, calling a function within the same function
// factorial
// example

func factorial(n int) int {

	n--

	if n == 0 {
		return 1
	}

	return factorial(n) * n

}

func main() {

	fmt.Println(factorial(10))

}
