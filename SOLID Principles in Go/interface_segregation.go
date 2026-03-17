package main

import "fmt"

// SOLID -> I
// Interface Segregation Principle
type Worker interface {
	Dowork()
}

type Eater interface {
	Eat()
}

type Robot struct{}
type Human struct{}

func (r Robot) Dowork() {
	fmt.Println("Robot is working!")
}

func (h Human) Dowork() {
	fmt.Println("Human is working")
}

func (h Human) Eat() {
	fmt.Println("Human is eating")
}

func Segregation() {
	r := Robot{}
	h := Human{}

	r.Dowork()

	h.Dowork()
	h.Eat()
}
