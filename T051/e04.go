package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age)
}

func main() {

	p0 := person{
		first: "Miss",
		last:  "Monyepenny",
		age:   34,
	}

	p0.speak()

}
