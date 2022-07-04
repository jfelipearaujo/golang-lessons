package main

import "fmt"

func main() {
	channel := make(chan int)

	go send(channel)

	receive(channel)
}

func send(s chan<- int) {
	s <- 42
	fmt.Println("Sending 42 to channel")
}

func receive(r <-chan int) {
	fmt.Println("Received", <-r, " to channel")
}
