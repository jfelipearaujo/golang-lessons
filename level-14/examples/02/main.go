package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmplName string) {
	tmpl, err := template.ParseFiles("./templates/" + tmplName)

	if err != nil {
		http.Error(w, "Error when try to parse the template file. Details: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "Error when try to read the template file. Details: "+err.Error(), http.StatusInternalServerError)
	}
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
