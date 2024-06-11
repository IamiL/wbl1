package main

import "fmt"

func Conveyor(values []int, ch1 chan int, ch2 chan int) {
	for _, v := range values {
		ch1 <- v
		ch2 <- v * 2
	}
	close(ch1)
	close(ch2)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch1, ch2 := make(chan int, len(arr)), make(chan int, len(arr))

	Conveyor(arr, ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}
}
