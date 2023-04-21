package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	p2 := person{
		First:   "Miss",
		Last:    "Moneypenny",
		Sayings: []string{"Hello darling.", "When are you taking me for a date?", "I ll get the minister."},
	}
	bs0, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs0))

	bs1, err := json.Marshal(p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs1))

}
