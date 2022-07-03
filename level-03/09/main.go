package main

import "fmt"

func main() {
	favoriteSports := "soccer"

	switch favoriteSports {
	case "soccer":
		fmt.Println("soccer")
	case "basketball":
		fmt.Println("basketball")
	case "volley":
		fmt.Println("volley")
	}
}
