/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
	circle area= π r 2
	square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/
package main

import (
	"fmt"
	"math"
)

type Square struct {
	height, width float64
}

func (s Square) area() float64 {
	return s.height * s.width
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * (c.radius * 2)
}

type Shape interface {
	area() float64
}

func info(s Shape) float64 {
	return s.area()
}

func main() {
	s1 := Square{
		height: 44,
		width:  10,
	}

	c1 := Circle{
		radius: 33,
	}

	// Printing the area with the methods attached to each struct
	fmt.Println("Printing the area with the methods attached to each struct:")
	fmt.Println(s1.area())
	fmt.Println(c1.area())

	// Printing the area with the function that accepts shape interface
	fmt.Println("Printing the area with the function that accepts shape interface:")
	fmt.Println(info(s1))
	fmt.Println(info(c1))
}
