package main

import (
	"fmt"
)

// using iota to increment before bitshift

const (
	_  = iota
	kb = 1 << (iota * 10) // kb := 1024
	mb = 1 << (iota * 10) // mb := 1024 * kb
	gb = 1 << (iota * 10) // gb := 1024 * mb
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
