package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	counter := 0
	totalOrGoRoutines := 1000

	var waitGroup = new(sync.WaitGroup)
	waitGroup.Add(totalOrGoRoutines)

	var mutex = new(sync.Mutex)

	for i := 0; i < totalOrGoRoutines; i++ {
		go func() {
			// Trava o acesso ao contador
			mutex.Lock()

			number := counter
			// Informa que o processador pode executar outra GoROUTINE/thread
			runtime.Gosched()
			number++
			counter = number

			// Libera o acesso ao contador
			mutex.Unlock()

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println()

	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)

	// Para analisar se o programa tem condições de corrida (race conditions)
	// Caso tenha, o código NÃO é seguro
	// go run -race main.go
}
