package main

import "fmt"

// variable types

// `` used when a string is on multiple lines
// (raw string literals will keep spaces)

var t = 1000
var c = "test"
var e = false
var c01 = `test
test
test`

var e00 int
var e01 float64
var e10 bool
var e11 string

func main() {

	fmt.Printf("Variabila %v are tipul %T\n", t, t)
	fmt.Printf("Variabila %v are tipul %T\n", c, c)
	fmt.Printf("Variabila %v are tipul %T\n", c01, c01)
	fmt.Printf("Variabila %v are tipul %T\n", e, e)

	// the initialisation of zero value variables

	fmt.Printf("Variabila e00 de tipul %T are acum valoarea de initializare: %v\n", e00, e00)
	fmt.Printf("Variabila e01 de tipul %T are acum valoarea de initializare: %v\n", e01, e01)
	fmt.Printf("Variabila e10 de tipul %T are acum valoarea de initializare: %v\n", e10, e10)
	fmt.Printf("Variabila e11 de tipul %T are acum valoarea de initializare: %v\n", e11, e11)

}
