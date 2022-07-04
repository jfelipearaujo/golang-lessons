package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitGroup = new(sync.WaitGroup)

func main() {
	// Informo que a função main terá 1 GoROUTINE
	waitGroup.Add(2)

	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	// A função abaixo se tornou uma GoROUTINE
	// A GoROUTINE é uma thread que executa em paralelo.
	// Neste código temos 2 threads
	// A primeira thread executa a função main()
	// A segunda thread executa a função func2().
	go func1()
	go func2()

	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	// Forço a função main esperar todas as GoROUTINEs terminarem
	waitGroup.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)

		time.Sleep(time.Millisecond * 10)
	}

	// Informo que a função func1 terminou
	waitGroup.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)

		time.Sleep(time.Millisecond * 20)
	}

	// Informo que a função func2 terminou
	waitGroup.Done()
}
