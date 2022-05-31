package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				// ops++ it will give different value if the variable is not adding in atomic way

			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("atomic unsigned int value:", ops)
}
