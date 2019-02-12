package main

import "fmt"

type rect struct {
	w, h float32
}

func (r *rect) area() int {
	return int(r.h * r.w)
}

func (r rect) perim() int {
	return int(2*r.w + 2*r.h)
}

func myArea(r *rect) int {
	return int(r.h * r.w)
}

func main() {

	r := rect{w: 10, h: 20}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())

	fmt.Println(myArea(rp))
}
