package main

import "fmt"

func main() {
	fmt.Println("number:", number())

	num, txt := numberAndText()

	fmt.Println("number & text:", num, txt)
}

func number() int {
	return 0
}

func numberAndText() (int, string) {
	return 1, "Hello"
}
