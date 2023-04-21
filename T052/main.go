package main

import "fmt"

func main() {

	a := 12

	fmt.Println(a)
	fmt.Println(&a)

	// * gives you the adress

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// & gives you the value stored at an adress when you have the adress

	// can t assign *int to int because they are of different type
	var b *int = &a
	fmt.Println(b)
	// derefrence the adress ... give me the value stored at that adress
	fmt.Println(*b)
	// the variable can be directly derefrenced
	fmt.Println(*&a)

	c := 14
	fmt.Println(&c)

	// the value at that adress
	*b = 1200

	fmt.Println(a)

}
