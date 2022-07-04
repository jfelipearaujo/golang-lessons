package main

import "fmt"

func main() {
	channel := make(chan int) // Canal bidirecional

	go doSomething(10, channel)

	printChannel(channel)
}

func doSomething(total int, channel chan<- int) {
	for i := 0; i < total; i++ {
		channel <- i // Envia o valor para o canal
	}

	close(channel) // Fecha o canal
}

func printChannel(channel <-chan int) {
	for v := range channel {
		fmt.Println(v)
	}
}
