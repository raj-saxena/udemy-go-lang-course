package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sidelength float64
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{5.0}
	t := triangle{base: 5.0, height: 7.5}
	printArea(s)
	printArea(t)
}
