package main

import "fmt"

func main() {

	t0 := test0()
	fmt.Println(t0)

	t1 := func() int {
		return 1200
	}()

	fmt.Println(t1)
	fmt.Printf("%T\n", t1)

	t01 := test1()
	fmt.Printf("%T\n", t01)

	t11 := t01()
	fmt.Println(t11)
	fmt.Printf("%T\n", t11)

	// this prints the function that is return by the function
	fmt.Println(test1()())

}

// returning a string

func test0() string {
	s := "test01"
	return s
}

func test1() func() int {
	return func() int {
		return 12
	}
}
