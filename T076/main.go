package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		// check the error
		fmt.Println(err)
		// return so that it will stop the code at this line
		return
	}

	// defer the closing of the file
	defer f.Close()
	// byte slice and err
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	// convert byte slice to string
	fmt.Println(string(bs))
}
