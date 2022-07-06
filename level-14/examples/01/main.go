package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an example web page for Go!")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server is listening on port", portNumber)

	err := http.ListenAndServe(portNumber, nil)

	if err != nil {
		fmt.Println("An error happened while the server has been started:", err)
	}
}
