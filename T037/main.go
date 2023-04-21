package main

import "fmt"

/*
type person struct {
	first string
	last  string
	age   int
}
*/

func main() {

	// anonymous struct litteral replace person directly
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   34,
	}

	fmt.Println(p1)

}
