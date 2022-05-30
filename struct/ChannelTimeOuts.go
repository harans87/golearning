package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result"
	}()

	select {
	case msg := <-c1:
		fmt.Println("done without timeout", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	select {
	case msg2 := <-c1:
		fmt.Println("done without timeout", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("this will not print")
	}
}
