package main

import (
	"fmt"
	"time"
)

func main() {
	const numOfJobs = 5
	// worker pools
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= numOfJobs; w++ {
		go worker(w, jobs, results)
	}

	//send data to channels
	for k := 1; k <= 10; k++ {
		jobs <- k
	}
	close(jobs)

	// receive data
	for l := 1; l <= 10; l++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 3
	}
}
