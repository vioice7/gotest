package main

import "fmt"

// break continue

func main() {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue // continues the loop without running the rest of the code
		}

		if i > 30 {
			break // brakes the loop
		}
		fmt.Println(i)
	}

}
