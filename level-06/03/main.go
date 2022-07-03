package main

import "fmt"

func main() {
	defer funcA()
	funcB()
}

func funcA() {
	fmt.Println("A: This will print at the end")
}

func funcB() {
	fmt.Println("B: This will print first")
}
