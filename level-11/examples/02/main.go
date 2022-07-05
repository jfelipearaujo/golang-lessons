package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := strings.NewReader("Hello, World!")

	io.Copy(file, reader)
}
