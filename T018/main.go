package main

import "fmt"

func main() {

	// print utf-8 characters 33 122

	for i := 33; i <= 122; i++ {

		fmt.Printf("number: %v\tcharacter %#U\t%c\n", i, i, i)

	}

}
