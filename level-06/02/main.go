package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println("A: total sum:", numberA(x...))
	fmt.Println("B: total sum:", numberB(x))
}

// As reticencias são chamadas de 'parâmetro variádico'
func numberA(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

func numberB(numbers []int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}
