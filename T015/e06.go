package main

import "fmt"

const (
	y0 = 2018
	y1 = y0 + iota
	y2 = y0 + iota
	y3 = y0 + iota
)

func main() {
	fmt.Println(y0)
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
}
