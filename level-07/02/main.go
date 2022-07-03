package main

import "fmt"

type person struct {
	name string
	age  int
}

// Espero receber um ponteiro de memória
func changeMe(p *person) {
	p.name = "Todd"
	p.age++
}

func main() {
	p := person{
		name: "John",
		age:  30,
	}

	fmt.Println(p)

	// Aqui eu informo o endereço de memória da struct
	changeMe(&p)

	fmt.Println(p)
}
