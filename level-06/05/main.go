package main

import (
	"fmt"
	"math"
)

// -------------------------------------------------
type Quadrado struct {
	lado float64
}

func (q Quadrado) area() float64 {
	return q.lado * q.lado
}

// -------------------------------------------------

// -------------------------------------------------
type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

// -------------------------------------------------

type info interface {
	area() float64
}

func measure(i info) {
	fmt.Printf("The area of a %T is %v\n", i, i.area())
}

func main() {
	q := Quadrado{
		lado: 10,
	}
	c := Circulo{
		raio: 10,
	}

	measure(q)
	measure(c)
}
