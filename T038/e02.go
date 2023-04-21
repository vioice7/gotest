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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for j, va := range v.favorite_ice_cream_flavors {
			fmt.Println(j)
			fmt.Println(va)
		}
		fmt.Println("")
	}

}
