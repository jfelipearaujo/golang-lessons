package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")

	if err != nil {
		fmt.Println("err happened ->", err)
		return
	}

	defer file.Close()

	log.SetOutput(file)

	noFile, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("err happened ->", err)
	}

	defer noFile.Close()

	fmt.Println("Check the log.txt file in the directory")
}
