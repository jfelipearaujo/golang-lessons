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
	_, err := sqrt(-10)

	if err != nil {
		log.Println(err)
	}
}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("math: square root of negative number: %v", i),
		}
	}

	return 42, nil
}
