package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	start := time.Now()

	Sleep(time.Second * 2)

	fmt.Println("completion of work", time.Since(start))
}
