package main

import "fmt"

func main() {

	// deleting from a slice

	x := []int{11, 23, 45, 67, 89, 90}
	fmt.Println(x)
	x = append(x, 12, 14, 16, 18, 20)
	fmt.Println(x)

	y := []int{100, 200, 300}
	x = append(x, y...)
	fmt.Println(x)

	// append some slices: a slice till position 2 and from position 4
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
