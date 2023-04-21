package main

import "fmt"

func main() {

	c := 0
	for i := 65; i <= 90; i++ {

		c++
		fmt.Printf("%v\n", c)

		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}

	}

}
