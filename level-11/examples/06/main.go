package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")

	if err != nil {
		// O log Fatal chama os.Exit(1)
		// isso significa que o programa termina com um código de erro de 1
		// e o defer não é executado
		log.Fatalln(err)
	}

	// do nothing
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}
