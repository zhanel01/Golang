package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Define the Circle struct
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Define the Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to print the area of a Shape
func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	PrintArea(circle)
	PrintArea(rectangle)
}
