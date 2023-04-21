package main

import "fmt"

// struct is a composite data type (aggregate data type)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   34,
	}

	p2 := person{
		first: "Miss",
		last:  "Monyepenny",
		age:   30,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
