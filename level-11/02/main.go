package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Any last wishes?",
			"Never say never",
		},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}

func toJSON(p1 person) ([]byte, error) {
	bs, err := json.Marshal(p1)

	if err != nil {
		return []byte{}, fmt.Errorf("error when try to convert to JSON: %v", err)
	}

	return bs, nil
}
