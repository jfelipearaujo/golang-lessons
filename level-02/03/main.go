package main

import "fmt"

const (
	x = 1
	y = 2
)

const (
	a int = 1
	b int = 2
)

func main() {
	fmt.Println("Typed consts")
	fmt.Println("x :", x)
	fmt.Println("y :", y)

	fmt.Println("Non-Typed consts")
	fmt.Println("a :", a)
	fmt.Println("b :", b)
}
