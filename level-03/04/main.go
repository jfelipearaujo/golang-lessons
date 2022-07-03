package main

import "fmt"

func main() {
	i := 1993

	for {
		if i <= 2022 {
			fmt.Println(i)
			i++

			continue
		}

		break
	}
}
