package main

import "fmt"

func main() {

	bd := 1986

	for bd >= 1986 {
		fmt.Println(bd)
		if bd >= 2021 {
			break
		}
		bd++

	}

}
