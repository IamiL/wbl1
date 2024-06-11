package main

import (
	"fmt"
	"math/big"
)

// multiplication
func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// division
func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

// addition
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// subtraction
func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	a, b := big.NewInt(14000000), big.NewInt(7000000)

	fmt.Println("multiplication: ", multiply(a, b))
	fmt.Println("division: ", divide(a, b))
	fmt.Println("addition: ", add(a, b))
	fmt.Println("subtraction: ", sub(a, b))

}
