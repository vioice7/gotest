package main

import "fmt"

func main() {

	if true {
		fmt.Println("1")
	}

	if false {
		fmt.Println("0")
	}

	if !true {
		fmt.Println("00")
	}

	if !false {
		fmt.Println("11")
	}

	if x := 10; x == 10 {
		// x has only a scope here
		fmt.Printf("test01")
	}

	// multiple statements on a line:
	// fmt.Printf("test0"); fmt.Printf("test1")

}
