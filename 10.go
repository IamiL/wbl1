package main

import "fmt"

func main() {
	values := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	sorted := make(map[int][]float32)
	for _, v := range values {
		sorted[int(v)/10*10] = append(sorted[int(v)/10*10], v)
	}

	fmt.Println(sorted)
}
