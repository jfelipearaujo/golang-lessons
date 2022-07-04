package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var waitGroup sync.WaitGroup
var counter int32

func main() {
	goRoutines := 100

	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	genGoRoutine(goRoutines)

	waitGroup.Wait()

	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	fmt.Println("Counter:", atomic.LoadInt32(&counter))
}

func genGoRoutine(numOfGoRoutines int) {
	waitGroup.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)

			runtime.Gosched()

			waitGroup.Done()
		}()
	}
}
