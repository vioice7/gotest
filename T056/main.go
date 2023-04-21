package main

import (
	"encoding/json"
	"fmt"
)

//

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":34},{"First":"Miss","Last":"Moneypenny","Age":30}]`
	// convert string into a slice of bytes
	byteslice := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", byteslice)
	// people is a slice of person
	//people := []person{}
	var people []person
	// error if is defined it should only assigned
	err := json.Unmarshal(byteslice, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)
	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
