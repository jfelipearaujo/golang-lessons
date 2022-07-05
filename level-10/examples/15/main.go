package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancela o contexto assim que o programa terminar

	for number := range gen(ctx) {
		fmt.Println("Work Number:", number)
		if number == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	number := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- number:
				number++
			}
		}
	}()

	return dst
}
