package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	var m = make(map[string]struct{})

	for _, v := range strings {
		m[v] = struct{}{}
	}

	fmt.Println(m)
}
