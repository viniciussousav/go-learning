package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func printArea(s Shape) {
	fmt.Printf("Area do quadrado %.2f \n", s.area())
}

type Square struct {
	length float64
	width  float64
}

func (s Square) area() float64 {
	return s.length * s.width
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func main() {
	square := Square{12, 12}
	printArea(square)

	circle := Circle{10}
	printArea(circle)
}
