package main

import "fmt"

func intersection(s1, s2 []int) []int {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	res := make([]int, 0)
	for _, e := range s2 {
		// if elements present in the hashmap then append intersection list.
		if hash[e] {
			res = append(res, e)
		}
	}

	return removeDupls(res)
}

// remove duplicates from slice.
func removeDupls(elements []int) []int {
	encountered := make(map[int]bool)
	res := make([]int, 0, len(elements))
	for _, element := range elements {
		if !encountered[element] {
			res = append(res, element)
			encountered[element] = true
		}
	}
	return res
}

func main() {
	slice1 := []int{7, 3, 9, 28, 1, 25, 15}
	slice2 := []int{15, 8, 5, 3, 1, 26, 14}

	a := intersection(slice1, slice2)

	fmt.Println(a)
}
