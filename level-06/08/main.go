package main

import "fmt"

func main() {
	x := 1
	y := func() func(input int) int {
		return func(input int) int {
			return input*2 + 1
		}
	}

	fmt.Println(y()(x))
}
