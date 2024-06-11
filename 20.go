package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		runes := []rune(input)
		runes = runes[:len(runes)-2]
		words := strings.Split(string(runes), " ")
		for keyLeft, keyRight := 0, len(words)-1; keyLeft < keyRight; keyLeft, keyRight = keyLeft+1, keyRight-1 {

			words[keyLeft], words[keyRight] = words[keyRight], words[keyLeft]
		}

		println(strings.Join(words, " "))
	}
}
