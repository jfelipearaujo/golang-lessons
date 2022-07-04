package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First_name string `json:"first_name"`
	Age        int    `json:"age"`
}

func main() {
	u1 := User{
		First_name: "James",
		Age:        32,
	}
	u2 := User{
		First_name: "Moneypenny",
		Age:        27,
	}
	u3 := User{
		First_name: "M",
		Age:        54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	jsonBytes, error := json.Marshal(users)

	if error != nil {
		fmt.Println("Error:", error)
	}

	fmt.Println(string(jsonBytes))
}
