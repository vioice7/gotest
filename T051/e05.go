package main

import (
	"fmt"
	"math"
)

type square struct {
	lenght float64
}

type circle struct {
	radius float64
}

type shape interface {
	// remember to
	// add the return
	// type to a
	// function
	area() float64
}

func (s square) area() float64 {
	return s.lenght * s.lenght
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	s0 := square{
		lenght: 211,
	}

	c0 := circle{
		radius: 12.34,
	}
	info(s0)
	info(c0)
}
