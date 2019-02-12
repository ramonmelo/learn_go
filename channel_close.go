package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("Received job", j)
				time.Sleep(time.Second)
			} else {
				fmt.Println("All jobs received")
				done <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		// time.Sleep(500 * time.Millisecond)
		jobs <- i
		fmt.Println("Send job", i)
	}
	close(jobs)

	fmt.Println("Wait done")
	<-done
}
