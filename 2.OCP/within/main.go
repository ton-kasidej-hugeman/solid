package main

import "fmt"

// Shape interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// PrintArea function operates on the Shape interface
func PrintArea(s Shape) {
	fmt.Println("Shape Area:", s.Area())
}

// Triangle New shape implementation
type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}
	t := Triangle{Base: 10, Height: 5}

	PrintArea(r)
	PrintArea(c)
	PrintArea(t)
}
