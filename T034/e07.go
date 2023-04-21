package main

import "fmt"

func main() {

	c := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooo, James"}}

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Println(c[i][j])
		}
	}

}
