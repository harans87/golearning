package main

import (
	"fmt"
	"time"
)

func main() {
	tick1 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tick1.C:
				fmt.Println("ticking at:", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	tick1.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
