package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("There was an error when square rooting a negative number.")
		// e := errors.New("There was an error when square rooting a negative number.")
		// fmt.Errorf allows us to pass a value to the string
		e := fmt.Errorf("There was an error when square rooting a negative number: %v.", f)
		return 0, sqrtError{
			lat:  "46 N",
			long: "25 E",
			err:  e,
		}
	}
	return 42, nil
}
