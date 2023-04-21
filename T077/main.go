package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error happened", err)

		// log prints by default to the
		// srandard io but can also be modified
		// also adds a time stamp
		log.Println("error happened", err)
	}

}
