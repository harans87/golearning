package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for j := 0; j < 3; j++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for r := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- r
		}
	}()

	burstyRequests := make(chan int, 5)
	for k := 1; k < 5; k++ {
		burstyRequests <- k
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
