package main

import "fmt"

func main() {
	val1, val2 := 1, 2
	val2, val1 = val1, val2

	fmt.Println(val1, val2)
}
