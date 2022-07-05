package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error Check 1:", ctx.Err())
	fmt.Println("Num of GoRoutines:", runtime.NumGoroutine())

	go func() {
		number := 0
		for {
			select {
			// O Done será recebido pelo contexto
			// quando o contexto for cancelado através do método cancel()
			case <-ctx.Done():
				return

			// Enquanto não receber o Done (cancel)
			// a goroutine vai incrementando o número, ou executando um trabalho
			default:
				number++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Work Number:", number)
			}
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("Cancelling context...")

	cancel()

	fmt.Println("Error Check 2:", ctx.Err())
	fmt.Println("Num of GoRoutines:", runtime.NumGoroutine())

	fmt.Println("Context cancelled")
}
