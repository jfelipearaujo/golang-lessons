package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Your OS is:", runtime.GOOS)
	fmt.Println("Your CPU is:", runtime.GOARCH)
}
