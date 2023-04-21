package main

import "fmt"

func test0() int {

	// a enclosed to this function (visible only here)

	a := 1000
	return a
}

func main() {

	fmt.Println(test0())

}
