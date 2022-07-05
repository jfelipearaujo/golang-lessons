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
		log.Panicln(err)
	}

	fmt.Println("This print is not in the panic")
}

func foo() {
	fmt.Println("When panic is called, deferred functions runs")
}
