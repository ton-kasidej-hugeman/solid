package main

import "fmt"

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

func PrintArea(shape interface{}) {
	switch s := shape.(type) {
	case Rectangle:
		fmt.Println("Rectangle Area:", s.Area())
	case Circle:
		fmt.Println("Circle Area:", s.Area())
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	PrintArea(r)
	PrintArea(c)
}
