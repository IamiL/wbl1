package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int32
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	atomic.AddInt32(&c.count, 1)
	wg.Done()
}

func main() {
	c := &Counter{count: 0}

	var wg sync.WaitGroup
	for i := 0; i < 999; i++ {
		wg.Add(1)
		go c.Increment(&wg)
	}

	wg.Wait()
	fmt.Println(c.count)
}
