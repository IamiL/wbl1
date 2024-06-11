package main

import "fmt"

func main() {
	slice := []string{"one", "two", "three"}
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}
