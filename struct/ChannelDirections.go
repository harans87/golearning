package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 2)

	// pings sends to pongs and pongs receives it and prints it.
	// pongs <- pings <- "receive only method"
	ping(pings, "receive only method")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
