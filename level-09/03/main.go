package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup
var counter int

func main() {
	goRoutines := 100

	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	genGoRoutine(goRoutines)

	waitGroup.Wait()

	fmt.Println("Num of GoROUTINES:", runtime.NumGoroutine())

	fmt.Println("Counter:", counter)
}

func genGoRoutine(numOfGoRoutines int) {
	waitGroup.Add(numOfGoRoutines)

	for i := 0; i < numOfGoRoutines; i++ {
		go func() {
			num := counter
			runtime.Gosched()
			num++
			counter = num

			waitGroup.Done()
		}()
	}
}
