package main

import "fmt"

func main() {
	x := 1

	switch {
	case x == 0:
		fmt.Println("0")
	case x == 1:
		fmt.Println("1")
	case x == 2:
		fmt.Println("2")
	}
}
