package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sidelength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return math.Pow(s.sidelength, 2)
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println(area)
}

func main() {
	t := triangle{
		base:   10,
		height: 10,
	}

	printArea(t)

	s := square{
		sidelength: 10,
	}

	printArea(s)
}
