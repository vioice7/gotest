package main

import (
	"fmt"
)

func main() {

	n := 40    // numbers of lines on a triangle
	level := 4 // numbers of levels
	blank := ""
	character := ""

	for t := 0; t < level; t++ {

		for i := 0; i < n; i++ {
			character = "-"
			for j := 1; j <= i; j++ {
				character = character + "-"
			}
			blank = ""
			for k := 0; k < n-i; k++ {
				blank += " "
			}
			// fmt.Println(blank + character + strconv.Itoa(i) + character)
			fmt.Println(blank + character + " " + character)
		}

		for i := n - 1; i >= 0; i-- {
			character = "-"
			for j := 1; j <= i; j++ {
				character = character + "-"
			}
			blank = ""
			for k := 0; k < n-i; k++ {
				blank += " "
			}
			// fmt.Println(blank + character + strconv.Itoa(i) + character)
			fmt.Println(blank + character + " " + character)
		}

	}

}
