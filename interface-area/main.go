package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		height: 10.0,
		base:   5.0,
	}

	squ := square{
		sideLength: 10,
	}

	printArea(tri)
	printArea(squ)
}

func printArea(sha shape) {
	fmt.Println(sha.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
