package main

import "fmt"

func main() {
	data := make(map[string]int)

	data["k1"] = 1
	data["k2"] = 8

	fmt.Println(data)
	fmt.Println(len(data))

	delete(data, "k2")

	fmt.Println(data)

	v2, prs := data["k2"]
	v1, prs2 := data["k1"]

	fmt.Println("Presence:", prs, v2)
	fmt.Println("Presence:", prs2, v1)

	fmt.Println(data["k3"])
}
