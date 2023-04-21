package main

import "fmt"

func main() {

	// maps composite literal
	m := map[string]int{
		"test0": 10,
		"test1": 11,
		"test2": 12,
		"test3": 21,
	}

	m["test00"] = 120

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("")
	// a non existing key can be deleted without error
	delete(m, "test00")

	for k, v := range m {
		fmt.Println(k, v)
	}
	// check if a map key can be deleted
	if v, ok := m["test3"]; ok {
		fmt.Println("ok ", v)
	}

}
