package main

import "fmt"

func main() {
	// create strings with "" for single line
	// create with `` for multiple lines
	s := `"Salutare acesta este un sir de caractere acuma."`

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// string is a slice of byte

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// utf8 representation
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("la pozitia %d avem caracterul in hexa %#x\n", i, v)
	}

}
