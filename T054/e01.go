package main

import "fmt"

func main() {

	// pointers and variables

	a := 10

	fmt.Println("The value of a is:", a)
	fmt.Println("The adress of a is:", &a)
	fmt.Println("The dereference adress *&a will give the value:", *&a)

}
