package main

import "fmt"

func main() {
	channel := make(chan int) // Canal bidirecional
	quit := make(chan int)

	go receive(50, channel, quit)

	send(channel, quit)
}

func receive(total int, channel chan int, quit chan int) {
	for i := 0; i < total; i++ {
		v := <-channel
		fmt.Println("Recebido:", v)
	}
	quit <- 0
}

func send(channel chan int, quit chan int) {
	a_number := 0

	for {
		select {
		case channel <- a_number:
			a_number++
		case <-quit:
			return
		}
	}
}
