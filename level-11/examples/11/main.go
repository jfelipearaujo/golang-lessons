package main

import (
	"fmt"
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
		return 0, fmt.Errorf("math: square root of negative number: %v", i)
	}

	return 42, nil
}
