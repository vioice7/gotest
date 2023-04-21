package main

import (
	"encoding/json"
	"fmt"
)

// the fields names should be upper case first

// JSON to go
// website
// transforming all

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   34,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   30,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	byteslice, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteslice))

}
