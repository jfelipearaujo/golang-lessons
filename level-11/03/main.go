package main

import "fmt"

type specialError struct {
	error string
}

func (se specialError) Error() string {
	return fmt.Sprintf("Special error: %v", se.error)
}

func main() {
	err := specialError{"Something went wrong"}

	doSomething(err)
}

func doSomething(e error) {
	fmt.Println("Error:", e.Error())
}
