package main

import "fmt"

func main() {
	channel := gen()

	receive(channel)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	channel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}

		close(channel)
	}()

	return channel
}

func receive(channel <-chan int) {
	for value := range channel {
		fmt.Println(value)
	}
}
