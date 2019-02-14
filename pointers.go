package main

import "fmt"

func zeroVal(i int) {
	i = 0
}

func zeroPtr(i *int) {
	*i = 0
}

func main() {
	var i = 10

	zeroVal(i)

	fmt.Println(i)

	zeroPtr(&i)

	fmt.Println(i)
}
