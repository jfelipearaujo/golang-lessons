package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	jsonData := []byte(
		`[
			{"name":"John","age":30,"occupation":"developer"},
			{"name":"Jane","age":25,"occupation":"designer"},
			{"name":"Jack","age":20,"occupation":"student"}
		]`)

	var people []Person

	err := json.Unmarshal(jsonData, &people)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("People:", people)
}
