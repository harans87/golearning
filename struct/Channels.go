package main

import (
	"fmt"
	"time"
)

func main() {
	// make a channel of string
	messages := make(chan string)

	go func() {
		messages <- "pingFromVSCode"
	}()

	msg := <-messages

	fmt.Println(msg)

	messages2 := make(chan string, 2)

	go another(messages2)

	time.Sleep(time.Second)

}

func another(ch chan string) {
	ch <- "anotherPing"
	msg2 := <-ch
	fmt.Println(msg2)
}
