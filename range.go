package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}
	sum := 0

	m := make(map[int]int)

	m[1] = 2
	m[2] = 3
	m[3] = 4

	for i, num := range s {
		fmt.Println(i)
		sum += num
	}

	fmt.Println(sum)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
