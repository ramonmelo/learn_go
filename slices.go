package main

import "fmt"

func main() {

	s := make([]string, 3)

	fmt.Println("Slice", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Slice", s)
	fmt.Println("Size", len(s))

	s = append(s, "d", "e")
	fmt.Println("Slice", s)

	c := make([]string, len(s))
	fmt.Println(c)
	copy(c, s)
	fmt.Println(c)

	fmt.Println(s[2:len(s)])
}
