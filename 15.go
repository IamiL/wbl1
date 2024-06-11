package main

import "math/rand"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// выделяем под justString необходимый объём памяти
	justString = string(append([]rune{}, []rune(v)[:100]...))
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, size)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}
