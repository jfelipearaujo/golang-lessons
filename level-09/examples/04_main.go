package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	var counter int32 = 0
	totalOrGoRoutines := 100

	var waitGroup = new(sync.WaitGroup)
	waitGroup.Add(totalOrGoRoutines)

	for i := 0; i < totalOrGoRoutines; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			runtime.Gosched()

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println()

	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())
	fmt.Println("Counter:", atomic.LoadInt32(&counter))

	// Para analisar se o programa tem condições de corrida (race conditions)
	// Caso tenha, o código NÃO é seguro
	// go run -race main.go
}
