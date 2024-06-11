package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	point1, point2 := New(8, 2), New(13, 27)
	d := math.Sqrt(math.Pow(point2.X()-point1.X(), 2) + math.Pow(point2.Y()-point1.Y(), 2))

	fmt.Println("The distance is equal to ", d)
}
