package main

import (
	"fmt"
)

type Timer struct {
	id    string
	value int
}

func (t *Timer) tick() int {

	t.value += 1

	return t.value
}

func main() {
	var x int
	fmt.Scanln(&x)

	t := Timer{"timer1", 0}

	for i := 0; i < x; i++ {
		fmt.Println(t.tick())
	}
}
