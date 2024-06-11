package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int8{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(value int8) {
			fmt.Printf("the square of the number %d:%d\n", value, value*value)
			wg.Done()
		}(numbers[i])
	}

	// waiting for the wait group to exit main on time
	wg.Wait()

}
