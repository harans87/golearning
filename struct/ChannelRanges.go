package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	for e := range queue {
		fmt.Println(e)
	}
}
