package main

import "fmt"

func main() {

	// maps composite literal
	m := map[string]int{
		"test0": 10,
		"test1": 11,
		"test2": 12,
		"test3": 13,
	}

	m["test00"] = 120

	for k, v := range m {
		fmt.Println(k, v)
	}

	c := []int{1, 5, 6, 7, 8, 9}

	for i, v := range c {
		fmt.Println(i, v)
	}

}
