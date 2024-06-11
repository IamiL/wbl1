/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от
родительской структуры Human (аналог наследования).*/

package main

import "time"

// Human declaration
type Human struct {
	firstName string
	lastName  string
	height    float32
	weight    float32
}

// Human constructor
func NewHuman(fName string, lName string, height float32, weight float32) *Human {
	return &Human{
		firstName: fName,
		lastName:  lName,
		height:    height,
		weight:    weight,
	}
}

// a method for weight gain
func (h *Human) IncreaseWeight(value float32) {
	h.weight = h.weight + value
}

// Action declaration
type Action struct {
	Human
	name     string
	duration time.Duration
}

// Action constructor with inheritance from Human
func NewAction(fName string, lName string, height float32, weight float32, actionName string, dur time.Duration) *Action {
	h := NewHuman(fName, lName, height, weight)
	return &Action{
		name:     actionName,
		duration: dur,
		Human:    *h,
	}
}

func main() {
	// initializing an Action
	act := NewAction("John", "Parker", 173.2, 76, "have lunch", time.Minute*5)
	// using the Human method
	act.IncreaseWeight(0.5)
}
