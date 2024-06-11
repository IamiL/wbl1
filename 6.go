package main

import (
	"context"
	"fmt"
)

func main() {
	// stopping using a channel
	stop := make(chan struct{})
	go func() {
		for i := 0; true; i++ {
			select {
			case <-stop:
				return

			default:
				fmt.Println(i)
			}
		}
	}()
	stop <- struct{}{}

	// stopping using context
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; true; i++ {
		select {
		case <-ctx.Done():

			return

		default:
			fmt.Println(i)
		}
	}
	cancel()

	// stopping using a flag variable
	var flag bool
	go func() {
		for i := 0; true; i++ {
			if flag {
				return
			} else {
				fmt.Println(i)
			}

		}
	}()
	flag = true
}
