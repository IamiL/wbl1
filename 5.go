package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	secondsValue, _ := strconv.Atoi(os.Args[1])

	// initialize the channel
	ch := make(chan string)

	// writing to the channel
	go func() {
		for i := 0; true; i++ {
			ch <- "data" + strconv.Itoa(i)
		}
	}()

	// reading from the channel
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	// exiting the function after the specified time has elapsed
	time.Sleep(time.Duration(secondsValue) * time.Millisecond * 1000)
}
