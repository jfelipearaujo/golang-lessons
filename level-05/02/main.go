package main

import "fmt"

type person struct {
	first_name        string
	last_name         string
	ice_cream_flavors []string
}

func main() {
	var people = make(map[string]person)

	p1 := person{
		first_name: "James",
		last_name:  "Bond",
		ice_cream_flavors: []string{
			"Chocolate",
			"Vanilla",
			"Strawberry",
		},
	}

	people[p1.last_name] = p1

	p2 := person{
		first_name: "Miss",
		last_name:  "Moneypenny",
		ice_cream_flavors: []string{
			"Vanilla",
			"Chocolate",
			"Strawberry",
		},
	}

	people[p2.last_name] = p2

	for _, person := range people {
		fmt.Println(person.first_name, person.last_name)
		for _, ice_cream_flavor := range person.ice_cream_flavors {
			fmt.Println(ice_cream_flavor)
		}
	}
}
