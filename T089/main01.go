package main

import "fmt"

func main() {

	var t string

	qs := make(chan bool)
	go TestGoRoutine(qs, &t)

	if <-qs == true {
		fmt.Println(t)
	}

	//fmt.Println(v)
}

func TestGoRoutine(qs chan bool, test *string) {
	*test = "this is a goroutine"
	qs <- false
}
