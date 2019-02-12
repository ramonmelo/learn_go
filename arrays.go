package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[2] = 10
	fmt.Println(a)
	fmt.Println(len(a))

	b := []int{1, 2, 3, 4, 10}

	fmt.Println(b)
}
