package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) NameAndAge() string {
	return fmt.Sprintf("%s is %d years old", p.name, p.age)
}

func main() {
	p := Person{
		name: "John",
		age:  30,
	}

	fmt.Println(p.NameAndAge())
}
