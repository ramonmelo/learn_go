package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	other := func(msg string) {
		for i := 0; i < 6; i++ {
			fmt.Println(msg, "-", i)
		}
	}

	go other("going")
	go f("indirect")

	fmt.Scanln()
	fmt.Println("done")
}
