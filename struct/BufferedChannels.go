package main

import "fmt"

func main() {
	messages := make(chan int, 3)

	messages <- 1
	messages <- 2
	messages <- 3

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
