package main

import (
	"fmt"
)

func main() {
	foo()

	fmt.Println("Returned normally from foo")
}

func foo() {
	defer func() {
		// Caso exista um panic, o defer é executado
		// e o valor do panic é capturado pelo recover
		// e o programa continua normalmente
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo", r)
		}
	}()

	fmt.Println("Calling bar")
	bar(0)
	fmt.Println("Returned normally from bar")
}

func bar(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		// Encerra a stack de execução informando o item que causou o panic
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Deferred in bar", i)
	fmt.Println("Printing in bar", i)
	bar(i + 1)
}
