package main

import (
	"fmt"
	"log"
)

func main() {
	var rrr string
	fmt.Println("Test panic system?")
	fmt.Scanf("%s", &rrr)
	if rrr == "yes" {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered sucessfully from panic.\n")
				// fmt.Printf("The panic error was: %v", r)
			}
		}()
		log.Panic("Panicking!")
	} else {
		log.Fatalln("Fatal Err: Testing some fatal error.")
	}
}
