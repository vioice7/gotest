package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	v0 := vehicle{
		doors: 4,
		color: "blue",
	}

	t0 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	s0 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(v0)
	fmt.Println("")
	fmt.Println(t0)
	fmt.Println("")
	fmt.Println(s0)
	fmt.Println("")

	fmt.Printf("The vehcile v0 has %v doors and is colored %v.\n", v0.doors, v0.color)
	fmt.Printf("The truck t0 has %v doors and is colored %v.\n", t0.doors, t0.color)
	fmt.Printf("The sedan s0 has %v doors and is colored %v.\n", s0.doors, s0.color)

}
