package main

import "fmt"

func main() {
	x := 1
	y := 2

	r := calculateSum(x, y, sum)

	fmt.Println(r)
}

func sum(a int, b int) int {
	return a + b
}

func calculateSum(a int, b int, sumFunc func(int, int) int) int {
	return sumFunc(a, b)
}
