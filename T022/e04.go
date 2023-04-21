package main

import "fmt"

func main() {

	bd := 1986

	for {
		if bd <= 2021 {
			fmt.Println(bd)
		} else {
			break
		}
		bd++

	}

}
