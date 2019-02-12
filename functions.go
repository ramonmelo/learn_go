package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 1, 2
}

func sums(nums ...int) int {
	fmt.Println(nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(plus(1, 2))

	a, b := vals()
	fmt.Println(a, b)

	fmt.Println(sums(1, 2, 3))

	values := []int{4, 5, 6}

	fmt.Println(sums(values...))

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = intSeq()

	fmt.Println(nextInt())
}
