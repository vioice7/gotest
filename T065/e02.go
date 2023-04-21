package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("test")
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p0 := person{
		name: "T0",
	}

	fmt.Println(p0)

	saySomething(&p0)
	//fmt.Println(p0.speak())

}
