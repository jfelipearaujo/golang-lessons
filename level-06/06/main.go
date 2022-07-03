package main

import "fmt"

func main() {
	x := 1
	y := func(input int) int {
		return input * 2
	}(x)

	fmt.Println(y)
}
