package main

import "fmt"

type vehicle struct {
	dors  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			dors:  2,
			color: "red",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			dors:  4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Println()

	fmt.Println(truck1.dors)
	fmt.Println(sedan1.color)
}
