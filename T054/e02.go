package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "Test0000"
	(*p).last = "Test1111"
	(*p).age = 21
}

func main() {

	p1 := person{
		first: "Test0",
		last:  "Test1",
		age:   12,
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}
