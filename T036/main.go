package main

import "fmt"

// struct is a composite data type (aggregate data type)
type person struct {
	first string
	last  string
	age   int
}

// embed person struct data type into secret agent datatype

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   34,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Monyepenny",
		age:   30,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println("")
	// sa1.person.first sa1.person.last is also more accurate
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
