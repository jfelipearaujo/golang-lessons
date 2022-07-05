package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (x mathError) Error() string {
	return fmt.Sprintf("a math error occurred (%v / %v) => %v", x.lat, x.long, x.err)
}

func main() {
	_, err := sqrt(-10)

	if err != nil {
		log.Println(err)
	}
}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		errMsg := fmt.Errorf("math: square root of negative number: %v", i)
		return 0, mathError{"50.2289 N", "99.4656 W", errMsg}
	}

	return 42, nil
}
