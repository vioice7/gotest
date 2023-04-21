package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// bcrypt to store a password
// generate hased password
// compare hashed password with a given passsword

func main() {

	s := "password123"
	byteslice, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(byteslice)
	logPwd := "password123"
	// hasshed password and the password (slice of bytes)
	err = bcrypt.CompareHashAndPassword(byteslice, []byte(logPwd))

	if err != nil {
		fmt.Println("The given password does not match the hashed pasword.")

	} else {
		fmt.Println("The givent password mathced the hashed password.")
	}

}
