package main

import (
	"fmt"
	"sync"
)

func main() {
	// in this example mutex is created to access counters from different go routines async.
	// share the counters state across multiple routines.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.increment(name)
		}
		wg.Done()
	}

	wg.Add(3)

	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()

	fmt.Println("Mutex output:", c.counters)
	fmt.Println("Mutex output:", c)

}

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
