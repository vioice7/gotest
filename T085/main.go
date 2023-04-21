package main

import "fmt"

func main() {

	fmt.Println("1 + 2 =", mySum(1, 2))
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 3 =", mySum(4, 3))
	fmt.Println("10 - 3 =", mySubstract(10, 3))

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func mySubstract(x int, y int) int {
	return x - y
}
