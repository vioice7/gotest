package main

import "fmt"

func main() {

	a := 10
	// switch no expression statement
	switch {
	case a < 10:
		fmt.Println("a0")
	case a == 10:
		fmt.Println("a10")
	case a > 10:
		fmt.Println("a1")
	}

}
