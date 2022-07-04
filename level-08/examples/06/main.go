package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

// O tipo ByAge implementa o interface sort.Interface paro
// o tipo Person que é um slice de pessoas
// e ordena o slice de pessoas por idade (by age)
type ByAge []Person

// Determina o tamanho do slice
func (a ByAge) Len() int {
	return len(a)
}

// Troca os elementos de um slice
func (a ByAge) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// Verifica se o elemento de um slice é maior que o elemento de outro slice
func (a ByAge) Less(i int, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []Person{
		{"Bob", 31, "software engineer"},
		{"Jane", 30, "software engineer"},
		{"John", 33, "software engineer"},
		{"Jill", 29, "software engineer"},
		{"Jack", 32, "software engineer"},
	}

	fmt.Println(people)
	// Converte o slice de pessoas em um slice de ByAge
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
