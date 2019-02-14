package main

import (
	"fmt"
	"sort"
)

func main() {

	str := []string{"b", "c", "a"}

	sort.Strings(str)
	fmt.Println(str)

	ints := []int{5, 9, 1}

	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)

	fmt.Println("Sorted", s)
}
