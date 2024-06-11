package main

import (
	"fmt"
	"sort"
	"strconv"
)

func binarySearch(arr []int, target int) (int, bool) {
	sort.Ints(arr)

	// middle
	var mid int

	min := 0

	high := len(arr) - 1

	for min <= high {
		mid = (min + high) / 2

		if mid == target {
			return mid, true
		}

		if mid > target {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}

	return 0, false
}

func main() {
	arr := []int{5, 1, 6, 8, 3, 2, 7, 5}

	nowNum, res := binarySearch(arr, 2)

	if res {
		fmt.Printf("element has been found - %v", strconv.Itoa(nowNum))
	} else {
		fmt.Println("element was not found")
	}

}
