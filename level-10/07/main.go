package main

import "fmt"

func main() {
	channel := make(chan int)
	numOfGoRoutines := 10

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				channel <- j
			}
		}()
	}

	for i := 0; i < numOfGoRoutines*10; i++ {
		fmt.Println(i, "\t", <-channel)
	}
}
