package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	length float64
	width float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main () {
	firstCircle := Circle {10}
	firstRectangle := Rectangle {2, 4.5}
	fmt.Println(firstCircle.Area())
	fmt.Println(firstRectangle.Area())
}

