package main

import (
	"bufio"
	"os"
)

func main() {
	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		runes := []rune(input)
		runes = runes[:len(runes)-1]
		for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		println(string(runes))
	}
}
