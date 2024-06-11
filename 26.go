package main

import (
	"fmt"
)

func main() {

	str := "abБcdefjро"

	var m = make(map[int]struct{})
	runes := []rune(str)

	var flag bool

	for _, v := range runes {
		code := int(v)
		if code > 64 && code < 91 || code > 1039 && code < 1072 {
			code = code + 32
		}
		if _, ok := m[code]; ok {
			flag = true
			break
		} else {
			m[code] = struct{}{}
		}

	}

	fmt.Println(!flag)
}
