package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	go func() { messages <- "ping!" }()

	messages <- "ping3!"
	go func() { messages <- "ping2!" }()

	// var msg string

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
