package main

import "fmt"

func main() {

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	r := 0

	var g string

	fmt.Scanf("%s", &g)

	// fmt.Printf("%s", g)

	results = append(results, g)

	// fmt.Println(results)

	for i := 0; i < len(results); i++ {
		if results[i] == "w" {
			r += 3
		} else if results[i] == "d" {
			r += 1
		}
	}

	fmt.Println(r)

}
