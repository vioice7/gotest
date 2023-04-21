package main

import "fmt"

// recursion

func main() {

	n := factorial(4)

	fmt.Println(n)

	n00 := factloop(4)

	fmt.Println(n00)

	n11 := factloop(4)

	fmt.Println(n11)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factloop(n int) int {

	m := 1
	for i := 1; i <= n; i++ {
		m *= i
	}

	return m
}

func factloop01(n int) int {

	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total

}
