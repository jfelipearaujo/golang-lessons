package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("output.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
