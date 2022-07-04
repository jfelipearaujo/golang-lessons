package main

import "fmt"

func main() {
	channel := make(chan int)    // Canal bidirecional
	channelR := make(<-chan int) // Canal de só leitura
	channelW := make(chan<- int) // Canal de só escrita

	fmt.Println("-----")
	fmt.Printf("Canal bidirecional '%T'\n", channel)
	fmt.Printf("Canal só de leitura '%T'\n", channelR)
	fmt.Printf("Canal só de escrita '%T'\n", channelW)

	fmt.Println("-----")
	fmt.Printf("Convertendo um Canal '%T' para um Canal '%T'\n", channel, (<-chan int)(channel))
	fmt.Printf("Convertendo um Canal '%T' para um Canal '%T'\n", channel, (chan<- int)(channel))
}
