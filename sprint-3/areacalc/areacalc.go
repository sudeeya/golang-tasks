package areacalc

import "strings"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	firstSide, secondSide float64
	shape                 string
}

func NewRectangle(firstSide, secondSide float64, shape string) *Rectangle {
	return &Rectangle{
		firstSide:  firstSide,
		secondSide: secondSide,
		shape:      shape,
	}
}

func (r Rectangle) Area() float64 {
	return r.firstSide * r.secondSide
}

func (r Rectangle) Type() string {
	return r.shape
}

type Circle struct {
	radius float64
	shape  string
}

func NewCircle(radius float64, shape string) *Circle {
	return &Circle{
		radius: radius,
		shape:  shape,
	}
}

func (c Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func (c Circle) Type() string {
	return c.shape
}

func AreaCalculator(figures []Shape) (string, float64) {
	var (
		shapes         = make([]string, len(figures))
		area   float64 = 0.0
	)
	for i, figure := range figures {
		shapes[i] = figure.Type()
		area += figure.Area()
	}
	return strings.Join(shapes, "-"), area
}
