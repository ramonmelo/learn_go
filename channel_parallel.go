package main

import (
	"fmt"
	"time"
)

func worker(id int, pipe chan int, done chan bool) {
	for {
		job, more := <-pipe

		if more {
			fmt.Println("INFO:: Worker", id, "Job", job)
			time.Sleep(time.Duration(job) * time.Second)
			fmt.Println("DONE:: Worker", id, "Job", job)
		} else {
			break
		}
	}

	done <- true
}

func main() {
	pipe := make(chan int)
	done := make(chan bool)

	workerTotal := 8
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < workerTotal; i++ {
		go worker(i, pipe, done)
	}

	for _, n := range nums {
		fmt.Println("Sending", n)
		pipe <- n
	}

	close(pipe)

	for i := 0; i < workerTotal; i++ {
		<-done
		fmt.Println("Done", i)
	}

	fmt.Println("Done")
}
