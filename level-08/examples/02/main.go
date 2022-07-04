package main

import (
	"encoding/json"
	"fmt"
)

// Para exportar para JSON (ou para outros pacotes)
// os Tipos e os Campos precisam estar com a 1ª letra maiúscula
// e para gerar um JSON com a anotação desejada utilizamos as cráses (`)
// para criar as tags e chegar ao resultado esperado
type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	p1 := Person{
		Name:       "John",
		Age:        30,
		Occupation: "developer",
	}

	p2 := Person{
		Name:       "Jane",
		Age:        25,
		Occupation: "designer",
	}

	p3 := Person{
		Name:       "Jack",
		Age:        20,
		Occupation: "student",
	}

	json_p1, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("P1 - Error:", err)
	}

	json_p2, err := json.Marshal(p2)

	if err != nil {
		fmt.Println("P2 - Error:", err)
	}

	json_p3, err := json.Marshal(p3)

	if err != nil {
		fmt.Println("P3 - Error:", err)
	}

	fmt.Println("P1:", string(json_p1))
	fmt.Println("P2:", string(json_p2))
	fmt.Println("P3:", string(json_p3))
}
