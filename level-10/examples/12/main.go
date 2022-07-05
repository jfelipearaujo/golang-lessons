package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channelA := make(chan int)
	channelB := make(chan int)

	go sendNumbersToChannel(10, channelA)
	go sendToOtherChannel(channelA, channelB)

	for v := range channelB {
		fmt.Println(v)
	}
}

func sendNumbersToChannel(total int, channel chan int) {
	for i := 0; i < total; i++ {
		channel <- i
	}

	close(channel)
}

func sendToOtherChannel(channelA chan int, channelB chan int) {
	var wg sync.WaitGroup

	for v := range channelA {
		wg.Add(1)

		go func(x int) {
			channelB <- work(x)

			wg.Done()
		}(v)
	}

	wg.Wait()
	close(channelB)
}

func work(number int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))

	return number * 2
}
