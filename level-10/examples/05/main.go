package main

import "fmt"

func main() {
	num := 500
	channelA := make(chan int) // Canal bidirecional
	channelB := make(chan int) // Canal bidirecional

	go func(x int) {
		for i := 0; i < x; i++ {
			channelA <- i // Envia o valor para o canal
		}
	}(num / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			channelB <- i // Envia o valor para o canal
		}
	}(num / 2)

	for i := 0; i < num; i++ {
		select {
		case v := <-channelA:
			fmt.Println("Canal A:", v)
		case v := <-channelB:
			fmt.Println("Canal B:", v)
		}
	}
}
