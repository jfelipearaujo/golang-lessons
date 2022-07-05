package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Print("fmt.Println ->", "\t")
		fmt.Println("err happened", err)

		fmt.Print("log.Println ->", "\t")
		log.Println("err happened", err)

		fmt.Print("log.Fatalln ->", "\t")
		log.Fatalln("err happened", err)

		fmt.Print("panic ->", "\t")
		panic(err)
	}

	// do nothing
}
