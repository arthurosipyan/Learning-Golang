package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	heigth float64
	base   float64
}

type sqaure struct {
	sideLength float64
}

func main() {
	t := triangle{heigth: 10, base: 10}
	s := sqaure{sideLength: 10}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.base * t.heigth * 0.5
}

func (s sqaure) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
