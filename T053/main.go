package main

import "fmt"

// always use pass by value

func main() {

	x := 0

	foo(x)
	fmt.Println(x)

	fmt.Println("----------")

	foo0(&x)
	fmt.Println(x)

}

func foo(y int) {

	fmt.Println(y)
	y = 12
	fmt.Println(y)

}

// change the value directly at that adress
func foo0(y *int) {

	fmt.Println(y)
	fmt.Println(*y)
	// the value at that adress is
	*y = 12
	fmt.Println(y)
	fmt.Println(*y)

}
