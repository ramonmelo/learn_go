package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
