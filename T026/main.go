package main

import "fmt"

func main() {

	x := []int{11, 23, 45, 67, 89, 90}
	fmt.Println(x[0])
	fmt.Println(x)
	fmt.Println(x[1:])  // slice from position 1
	fmt.Println(x[1:3]) // slice from position 1 to 3 but not including 3
	fmt.Println(x[:3])  // slice till position 3 but not including 3

}
