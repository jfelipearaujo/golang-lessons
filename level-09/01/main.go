package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)

	go printOn()
	go printTwo()

	waitGroup.Wait()
}

func printOn() {
	fmt.Println("Hello World from Print One")

	waitGroup.Done()
}

func printTwo() {
	fmt.Println("Hello World from Print Two")

	waitGroup.Done()
}
