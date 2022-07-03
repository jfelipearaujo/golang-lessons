package main

import "fmt"

func main() {

	// Map is a dictionary
	people := map[string][]string{
		"Silva_Fulano":    {"Jogar bola", "Games"},
		"Santos_Beltrano": {"Ler", "Cinema"},
	}

	for key, value := range people {
		fmt.Println(key)
		for i, v := range value {
			fmt.Printf("\t%d: %s\n", i, v)
		}
	}

	fmt.Println()

	people["Silva_Ciclano"] = []string{"Jogar bola", "Games", "Ler"}

	delete(people, "Silva_Fulano")

	for key, value := range people {
		fmt.Println(key)
		for i, v := range value {
			fmt.Printf("\t%d: %s\n", i, v)
		}
	}
}
