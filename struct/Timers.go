package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(2 * time.Second)

	<-t1.C
	fmt.Println("t1 has fired")

	t2 := time.NewTimer(time.Millisecond)

	go func() {
		<-t2.C
		fmt.Println("t2 has fired")
	}()
	stop1 := t2.Stop()
	if stop1 {
		fmt.Println("t2 has stopped")
	}
	time.Sleep(2 * time.Second)
}
