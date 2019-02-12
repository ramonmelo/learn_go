package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	w, h float64
}

type circle struct {
	r float64
}

// Methods declaration

func (r rect) area() float64 {
	return r.h * r.w
}

func (r rect) perim() float64 {
	return 2*r.h + 2*r.w
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func measure(g geometry) {
	fmt.Println(g.(type))
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {
	r := rect{5, 10}
	c := circle{4}

	measure(r)
	measure(c)
}
