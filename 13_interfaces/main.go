package main

import (
	"fmt"
	"math"
)

// interfaces: a set of method signatures for structs

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 2}
	rectangle := Rectangle{width: 3, height: 4}

	fmt.Printf("Circle Area = %+v\n", getArea(circle))
	fmt.Printf("Rectangle Area = %+v\n", getArea(rectangle))
}
