package main

type shape2D interface {
	area() float64
	perimeter() float64
}

type triangle struct {
	a float64
	b float64
	c float64
}

func (t triangle) area() float64 {
	return t.a * t.b * t.c
}

func (t triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
