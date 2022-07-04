package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channelA := work("A")
	channelB := work("B")
	channelC := converge(channelA, channelB)

	for i := 0; i < 10; i++ {
		fmt.Println(<-channelC)
	}
}

func work(message string) chan string {
	channel := make(chan string)

	go func(message string, channel chan string) {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Job %v say: %v", message, i)

			// Simbolizando uma tarefa pesada
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}(message, channel)

	return channel
}

func converge(channelA, channelB chan string) chan string {
	channel := make(chan string)

	go func() {
		for {
			// Recebe os valores do canal A e insere no canal resultante
			channel <- <-channelA
		}
	}()

	go func() {
		for {
			// Recebe os valores do canal B e insere no canal resultante
			channel <- <-channelB
		}
	}()

	return channel
}
