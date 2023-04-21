package main

import "fmt"

func main() {

	// switch statement (missing switch expression defaults to true)

	switch {
	case false:
		fmt.Println("nu o sa fie afisata")
	case true:
		fmt.Println("o sa fie afisata")
		fallthrough // not to be used
	case true:
		fmt.Println("asta o sa fie afisata")
	case false:
		fmt.Println("asta n-o sa fie afisata pt ca nu merge in continuare")
	default:
		fmt.Printf("asta e implicit")
	}

}
