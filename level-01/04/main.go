package main

import "fmt"

type jose int

func main() {
	var x jose

	fmt.Printf("x = %v\n", x)
	fmt.Printf("x = %T\n", x)

	x = 42

	fmt.Printf("x = %v\n", x)
}
