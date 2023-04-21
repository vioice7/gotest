package main

import "fmt"

type error interface {
	Error() string
}

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	// fmt.Println("foo ran -", e, "\n")

	// use assertion to acces e custom error (not conversion)
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}
