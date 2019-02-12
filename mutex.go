package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readOps uint64
	var writeOps uint64

	for i := 0; i < 100; i++ {
		go func() {
			total := 0

			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
}
