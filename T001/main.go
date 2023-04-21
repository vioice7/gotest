package main

import "fmt"

func main() {

	// variatic function (unlimited number of values of any type)
	// functin Println from package fmt
	fmt.Println("test", 10, true)

	// test the return of the Println function
	// n number of bytes, err error

	n, err := fmt.Println("ttttttttt", 1000, true, true, true)

	fmt.Println("n este: ", n)
	fmt.Println("eroarea este: ", err)

	// all variables must be used

	// _ caracter is used to throw away returns
	// err in our case is a variable that won't be used

	e, _ := fmt.Println("ttt")
	fmt.Println(e)

}
