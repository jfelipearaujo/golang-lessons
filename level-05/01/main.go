package main

import "fmt"

type person struct {
	first_name        string
	last_name         string
	ice_cream_flavors []string
}

func main() {
	p1 := person{
		first_name: "James",
		last_name:  "Bond",
		ice_cream_flavors: []string{
			"Chocolate",
			"Vanilla",
			"Strawberry",
		},
	}

	p2 := person{
		first_name: "Miss",
		last_name:  "Moneypenny",
		ice_cream_flavors: []string{
			"Vanilla",
			"Chocolate",
			"Strawberry",
		},
	}

	fmt.Println(p1.first_name, p1.last_name)
	for i, v := range p1.ice_cream_flavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first_name, p2.last_name)
	for i, v := range p2.ice_cream_flavors {
		fmt.Println(i, v)
	}

}
