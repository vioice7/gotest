package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		// check the error
		fmt.Println(err)
		// return so that it will stop the code at this line
		return
	}
	defer f.Close()

	r := strings.NewReader("Sal")
	io.Copy(f, r)
}
