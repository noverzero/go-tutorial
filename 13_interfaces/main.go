package main

import (
	"fmt"
	"math"
)

//Define Interface
type Shape interface {
	area() float64
}

type Circle struct {
	x      float64
	y      float64
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
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
	circle1 := Circle{0, 0, 5}
	rectangle1 := Rectangle{10, 5}

	fmt.Printf("circle1 Area: %f\n", getArea(circle1))
	fmt.Printf("rectangle1 Area: %f\n", getArea(rectangle1))

}
