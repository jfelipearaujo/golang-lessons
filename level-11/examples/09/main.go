package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, errors.New("math: square root of negative number")
	}

	return 42, nil
}
