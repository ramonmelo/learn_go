package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("start working")
	time.Sleep(time.Second * 2)
	fmt.Println("Done work")

	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)
	fmt.Println("Waiting...")

	<-done
}
