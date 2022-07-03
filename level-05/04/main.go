package main

import "fmt"

func main() {
	// Anonymous struct
	x := struct {
		name           string
		age            int
		hobbies        []string
		favorite_songs map[string]string
	}{
		name: "James",
		age:  32,
		hobbies: []string{
			"Guitar",
			"Basketball",
			"Coding",
		},
		favorite_songs: map[string]string{
			"Rock": "Thriller",
			"Pop":  "I Will Always Love You",
			"Jazz": "Symphony No. 5",
		},
	}

	fmt.Println(x)
}
