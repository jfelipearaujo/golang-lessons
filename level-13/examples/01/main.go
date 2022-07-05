package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(multiplication(1, 2, 3, 4, 5))
}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func multiplication(x ...int) int {
	total := 1

	for _, v := range x {
		total *= v
	}

	return total
}
