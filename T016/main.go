package main

import "fmt"

func main() {

	// for init;condition;post   {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("Outer loop: %v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tInner loop: %v\n", j)
		}

	}

	x := 1

	for x < 10 {
		fmt.Println(x)
		x++
	}

}
