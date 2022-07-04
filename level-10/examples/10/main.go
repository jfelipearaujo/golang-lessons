package main

import (
	"fmt"
	"sync"
)

func main() {
	odd := make(chan int)
	even := make(chan int)
	converge := make(chan int)

	go sendNumbers(odd, even)
	go receiveNumbers(odd, even, converge)

	for v := range converge {
		fmt.Println("Converge", v)
	}
}

func sendNumbers(odd chan int, even chan int) {
	total := 10

	for i := 0; i < total; i++ {
		if i%2 == 0 {
			odd <- i
		} else {
			even <- i
		}
	}

	close(odd)
	close(even)
}

func receiveNumbers(odd chan int, even chan int, converge chan int) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for v := range odd {
			converge <- v
		}

		wg.Done()
	}()

	wg.Add(1)

	go func() {
		for v := range even {
			converge <- v
		}

		wg.Done()
	}()

	wg.Wait()
	close(converge)
}
