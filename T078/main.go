package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// the log information will go to the file already made
	log.SetOutput(f)

	// checking the error for the next file and sending
	// the information to the log file
	f2, err := os.Open("no-file.txt")
	if err != nil {
		//
		log.Println("error happened", err)
	}
	defer f2.Close()

	fmt.Println("check log file in the directory")
}
