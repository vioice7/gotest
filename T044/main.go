package main

import "fmt"

// interfaces allow us to define behaviour and do polymorfism

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// a method (ataches the function to that type)

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, " - the secret agent speaking.")
}

// anybody who has this method speak() is also of type human

type human interface {
	speak()
}

// every value implements empty interface

func speak(h human) {

	// this is assertion
	switch h.(type) {
	case person:
		fmt.Println("I called human (person)", h.(person).first)
	case secretAgent:
		fmt.Println("I called human (secret agent)", h.(secretAgent).first)
	}
	fmt.Println("I called human", h)
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " - the person speaking.")
}

func main() {

	// this value is secretAgent and also human
	// (a value can be of more than one type)
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	// onlyfunction_speak can take type secretAgent (sa1 and sa2)
	// also this can take type person (p1)
	// this is called polymorfism because it cantake multiple types
	speak(sa1)
	speak(sa2)
	speak(p1)

	// this is conversion
	type hotdog int
	var x hotdog
	x = 10
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
