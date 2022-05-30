package main

import (
	"fmt"
	"time"
)

func worker(d chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done..")

	// once work is done send the signal
	d <- true
}

func main() {
	d := make(chan bool, 1)

	go worker(d)

	//this line makes the function blocking call, wait until done is sent.
	<-d

}
