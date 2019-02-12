package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan string, 1)

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"
	messages <- "hi2"

	select {
	case messages <- msg:
		fmt.Println("Message send", msg)
	default:
		fmt.Println("No message send")
	}

	select {
	case sig := <-signals:
		fmt.Println("received signal", sig)
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no activity")
	}

	fmt.Println(<-messages)
}
