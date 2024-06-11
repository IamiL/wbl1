package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func generate() interface{} {
	switch rand.Intn(4) {
	case 0:
		return 1

	case 1:
		return "1"

	case 2:
		return true

	case 3:
		return make(chan bool)

	}
	return 1
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(reflect.TypeOf(generate()))
	}
}
