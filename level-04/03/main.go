package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	first := slice[:3]

	fmt.Println("first:", first)

	second := slice[4:]

	fmt.Println("second:", second)

	third := slice[1:7]

	fmt.Println("third:", third)

	fourth := slice[2 : len(slice)-1]

	fmt.Println("fourth:", fourth)
}
