package main

import "fmt"

// anonymous struct for a vehicle

func main() {
	v0 := struct {
		milage int
		color  string
		parts  map[string]int
		oil    []string
	}{
		milage: 1000,
		color:  "blue",
		parts: map[string]int{
			"doors":      4,
			"wheels":     4,
			"headlights": 2,
			"taillights": 2,
		},
		oil: []string{
			"Havoline",
			"Mobile",
			"Shell",
		},
	}

	fmt.Println("Vehicle: ", v0)

	fmt.Println("")

	for j, v := range v0.parts {
		fmt.Println(j, v)
	}

}
