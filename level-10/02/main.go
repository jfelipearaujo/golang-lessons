package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
	}()

	fmt.Println(<-channel)

	fmt.Printf("------\n")
	fmt.Printf("c\t%T\n", channel)
}
