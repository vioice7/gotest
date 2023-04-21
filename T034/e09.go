package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["minister_m"] = []string{"Office", "Top secret", "Brandy"}

	for i, v := range m {
		for j, v0 := range v {
			fmt.Printf("%v %v, \t %v, %v \n", i, j, v, v0)
		}
	}

}
