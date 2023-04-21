package main

import (
	"fmt"
)

func main() {

	// maps composite literal
	m := map[string]int{
		"test0": 10,
		"test1": 11,
		"test2": 12,
		"test3": 13,
	}

	fmt.Println(m)
	fmt.Println(m["test0"])
	// this is not found but returns 0
	// fmt.Println(m["test"])

	v, ok := m["test"]

	fmt.Println(v)
	fmt.Println(ok)

	// check if the key is in the map

	if v, ok := m["test0"]; ok {
		fmt.Printf("key %v is in the map", v)
	}

}
