// this package should be in the GOPATH env variable
package main

import (
	"fmt"
)

// the sum of all numbers passed as variatic parameter
// also it will return an int
func mySum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}

func main() {
	fmt.Println("1 + 2 =", mySum(1, 2))
	fmt.Println("3 + 4 =", mySum(3, 4))
	fmt.Println("56 + 78 =", mySum(56, 78))
	fmt.Println("90 + 102 =", mySum(90, 102))
}
