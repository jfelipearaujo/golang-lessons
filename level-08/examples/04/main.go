package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	p := Person{
		Name:       "John",
		Age:        30,
		Occupation: "developer",
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(p)
}
