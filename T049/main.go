package main

import "fmt"

// closure to limit the scope of variables

var x int

// scope is on all package

func main() {

	var y int
	// scope is in all main function only

	{
		var z int

		// scope is only in this codeblock
		fmt.Println(z)

	}

	fmt.Println(x)

	fmt.Println(y)

	e0 := incrementor()
	e1 := incrementor()

	fmt.Println(e0())
	fmt.Println(e0())
	fmt.Println(e0())
	fmt.Println(e1())
	fmt.Println(e1())
	fmt.Println(e1())

}

func incrementor() func() int {
	var e int
	return func() int {
		e++
		return e
	}
}
