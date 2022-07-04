package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{name: "John", age: 30}

	// Abaixo é um atalho para (&p1).speak()
	p1.speak()

	// Abaixo é a "maneira mais correta"
	(&p1).speak()

	saySomething(&p1)
}
