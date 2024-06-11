package main

import (
	"fmt"
)

func main() {
	numbers := [5]uint8{2, 4, 6, 8, 10}

	var sum uint8

	// initialize the channel for transmitting the calculated squares of numbers
	ch := make(chan uint8, 5)

	for i := 0; i < 5; i++ {

		go func(value uint8, ch chan uint8) {
			// sending the calculated square of a number
			ch <- value * value
		}(numbers[i], ch)
	}

	// getting squares and changing the sum
	for i := 0; i < 5; i++ {
		sum = sum + <-ch
	}

	fmt.Println(sum)

}
