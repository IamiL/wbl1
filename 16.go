package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 4, 10, 7, 1}

	sort.Ints(arr)
	fmt.Println(arr)
}
