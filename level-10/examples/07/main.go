package main

import "fmt"

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan bool)

	go sendNumbers(odd, even, quit)

	receiveNumbers(odd, even, quit)
}

func sendNumbers(odd chan int, even chan int, quit chan bool) {
	total := 100

	for i := 0; i < total; i++ {
		if i%2 == 0 {
			odd <- i
		} else {
			even <- i
		}
	}

	close(odd)
	close(even)
	quit <- true
}

func receiveNumbers(odd chan int, even chan int, quit chan bool) {
	for {
		select {
		case v := <-odd:
			fmt.Println("Canal Par:", v)
		case v := <-even:
			fmt.Println("Canal Impar:", v)
		case <-quit:
			return
		}
	}
}
