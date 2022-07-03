package main

import "fmt"

func main() {
	a := calculate(1, 2)
	b := calculate(1, 2)

	fmt.Println("a:", a())
	fmt.Println("b:", b())
}

func sum(a int, b int) int {
	return a + b
}

func calculate(a int, b int) func() int {
	total := 0

	return func() int {
		total += sum(a, b)
		return total
	}
}
