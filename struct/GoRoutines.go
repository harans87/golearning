package main

import (
	"fmt"
	"time"
)

func f(name string) {
	for i := 0; i < 4; i++ {
		fmt.Println(name, ":", i)
	}
}

func main() {
	f("sync")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous")

	go f("async")

	time.Sleep(time.Second)

	fmt.Println("done")
}
