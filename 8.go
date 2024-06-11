package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := int64(1413412)

	binaryStr := []byte(strconv.FormatInt(value, 2))

	fmt.Println("value in binary system - ", string(binaryStr))

	//устанавливаем 2-й бит в 0
	binaryStr[2] = '0'

	newValue, _ := strconv.ParseInt(string(binaryStr), 2, 64)

	fmt.Println(newValue)
}
