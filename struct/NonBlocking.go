package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// msg := "hi"

	// go func() {
	// 	c1 <- "hi"
	// }()

	// time.Sleep(time.Second)

	select {
	case msg1 := <-c1:
		fmt.Println("msg1 received:", msg1)
	case msg2 := <-c2:
		fmt.Println("msg2 received", msg2)
	default:
		fmt.Println("none received")
	}
}
