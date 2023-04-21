package main

import "fmt"

var y = 911

func main() {

	// print directly to defaoult display

	fmt.Printf("Type of y is %T\n", y)

	fmt.Printf("The y in binary: %b\n", y)

	fmt.Printf("The y in hexa without 0x: %x\n", y)

	fmt.Printf("The y in hexa with 0x: %#x\n", y)

	// print to a string

	s := fmt.Sprintf("Value %v is %#x in hexa with 0x, %b in binary and %x in hexa without 0x.", y, y, y, y)

	// print the string s
	fmt.Println(s)
}
