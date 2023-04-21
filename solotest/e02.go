package main

import (
	"fmt"
)

func mars_age(age int) int {

	nzp := 365
	nzm := 687
	t := float64(nzp) / float64(nzm)
	ageo := float64(age)
	ageoo := t * ageo

	return int(ageoo)

}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
