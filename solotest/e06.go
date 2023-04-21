package main

import (
	"fmt"
)

func main() {

	var f int

	fmt.Printf("The frequency is: ")
	fmt.Scanln(&f)

	if f < 0 {
		fmt.Println("Wrong Input")
	} else if f < 20 {
		fmt.Println("Infrasound")
	} else if f >= 20 && f <= 20000 {
		fmt.Println("Audible")
	} else if f > 20000 {
		fmt.Println("Ultrasound")
	}

}
