package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	map_ := make(map[int]string)

	// allows you to solve the race condition problem
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1000; i >= 0; i-- {
			mutex.Lock()
			map_[i] = "data" + strconv.Itoa(i) + "recorded by gorutina 1"
			mutex.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 1000; i >= 0; i-- {
			mutex.Lock()
			map_[i] = "data" + strconv.Itoa(i) + "recorded by gorutina 2"
			mutex.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()

	for i := 0; i <= 1000; i++ {
		fmt.Println("key - ", i, ", value - ", map_[i])
	}
}
