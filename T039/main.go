package main

import "fmt"

//func ( r receiver ) identifier(parameters) (return(s)) { ... }
// when define a function you use parameters
// when you call a function you use arguments

func main() {
	test()
	test0("Aim")
	s0 := test1("Test01")
	fmt.Println(s0)
	x, y := test01("Aim0", "Aim1")
	fmt.Println(x)
	fmt.Println(y)
}

func test() {
	fmt.Println("sal")
}

// everything in Go is passed by value
func test0(s string) {
	fmt.Println("Sal", s)
}

func test1(st string) string {
	return fmt.Sprint("Hello from test, ", st)
}

func test01(t0 string, t1 string) (string, bool) {
	a := fmt.Sprint(t0, " ", t1)
	b := true
	return a, b
}
