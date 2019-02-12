package main

import "fmt"

func main() {

	queue := make(chan string, 3)

	queue <- "message 1"
	queue <- "message 2"
	queue <- "message 3"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
