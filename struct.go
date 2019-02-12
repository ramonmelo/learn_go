package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Lisa", age: 10})

	s := person{
		name: "James",
		age:  25,
	}

	sp := &s

	fmt.Println(s.name)
	fmt.Println(s.age)

	fmt.Println("Novo")
	fmt.Println(sp.name)
	fmt.Println(sp.age)
}
