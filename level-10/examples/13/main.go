package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channelA := make(chan int)
	channelB := make(chan int)
	maxNumberOfGoRoutines := 5

	go sendNumbersToChannel(100, channelA)
	go sendToOtherChannel(maxNumberOfGoRoutines, channelA, channelB)

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

func sendToOtherChannel(maxNumberOfGoRoutines int, channelA chan int, channelB chan int) {
	var wg sync.WaitGroup

	for i := 0; i < maxNumberOfGoRoutines; i++ {
		wg.Add(1)

		go func() {
			for v := range channelA {
				channelB <- work(v)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	close(channelB)
}

func work(number int) int {
	time.Sleep(time.Millisecond * 1000)

	return number
}
