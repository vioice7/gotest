package main

import "fmt"

func main() {

	// numeral coding system
	s := "E"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%x\n", n)
}
