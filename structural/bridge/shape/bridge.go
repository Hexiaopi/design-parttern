package shape

import "fmt"

// Implementor
type Color interface {
	GetColor() string
}

// Concrete Implementor
type Red struct{}

func (r Red) GetColor() string {
	return "red"
}

// Concrete Implementor
type Blue struct{}

func (b Blue) GetColor() string {
	return "blue"
}

// Abstraction
type Shape interface {
	Draw()
}

// Refined Abstraction
type Circle struct {
	color Color //组合关系
}

func (c Circle) Draw() {
	fmt.Println(c.color.GetColor() + " circle")
}

// Refined Abstraction
type Square struct {
	color Color //组合关系
}

func (s Square) Draw() {
	fmt.Println(s.color.GetColor() + " square")
}
