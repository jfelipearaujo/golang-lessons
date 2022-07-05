package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNegativeSqrt = errors.New("math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNegativeSqrt)

	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, ErrNegativeSqrt
	}

	return 42, nil
}
