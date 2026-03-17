package main

import "fmt"

// SOLID -> L
// Liskov Substituition Principle
type Shape interface {
	Area() float64
}
type Rect struct {
	length  float64
	breadth float64
}

func (r Rect) Area() float64 {
	return r.length * r.breadth
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	rect := Rect{length: 12, breadth: 6.28}
	square := Square{side: 4}
	PrintArea(rect)
	PrintArea(square)
}
