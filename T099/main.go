package main

import "fmt"

func main() {

	i := 0

	loop(&i)

	fmt.Println(i)
}

func loop(i *int) {

Test:
	if *i < 10 {
		*i++
		if *i > 5 {
			// fmt.Println()
		}
		fmt.Println(*i)
		goto Test
	}
}
