package main

import "fmt"

func main() {

	names := [][]string{
		{"Fulano", "da Silva", "Jogar boa"},
		{"Beltrano", "Santos", "Jogar games"},
		{"Siclano", "Silva", "Ler"},
	}

	for _, name := range names {
		for _, v := range name {
			fmt.Printf("%v\t", v)
		}
		fmt.Println()
	}
}
