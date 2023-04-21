package main

import "fmt"

func main() {

	n, err := fmt.Println("Sal")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}
