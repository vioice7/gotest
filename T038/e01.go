package main

import "fmt"

type person struct {
	first_name                 string
	last_name                  string
	favorite_ice_cream_flavors []string
}

func main() {
	p1 := person{
		first_name:                 "James",
		last_name:                  "Bond",
		favorite_ice_cream_flavors: []string{"chocolate", "strawberry", "martipan"},
	}
	p2 := person{
		first_name:                 "Miss",
		last_name:                  "Moneypenny",
		favorite_ice_cream_flavors: []string{"cherry", "vanilla", "martipan", "caju"},
	}

	fmt.Println(p1)
	fmt.Println("")
	fmt.Println(p2)
	fmt.Println("")

	for i, v := range p1.favorite_ice_cream_flavors {
		fmt.Println(i, v)
	}

	fmt.Println("")

	for i, v := range p2.favorite_ice_cream_flavors {
		fmt.Println(i, v)
	}

}
