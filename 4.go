package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func RecordRandomStr(ch chan string) {
	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	ch <- string(b)
	time.Sleep(time.Second)
}

func Worker(ctx context.Context, wg *sync.WaitGroup, input chan string) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return

		default:
			fmt.Println(<-input)
		}
	}
}

func main() {
	// getting the number of workers from command line arguments
	numberWorkers, _ := strconv.Atoi(os.Args[1])

	inputCh := make(chan string)

	// sending random strings to the channel
	go func() {
		for {
			RecordRandomStr(inputCh)
		}
	}()

	// контекст позволяет обработать и контролировать остановку воркеров
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}

	for i := 0; i < numberWorkers; i++ {
		workCtx, _ := context.WithCancel(ctx)
		wg.Add(1)
		go Worker(workCtx, wg, inputCh)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	cancel()

	fmt.Println("waiting for the workers to stop")
	wg.Wait()
}
