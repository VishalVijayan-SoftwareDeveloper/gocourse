package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width, length float64
}
type Circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.length
}
func (r rect) perimeter() float64 {
	return 2 * (r.width + r.length)
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {

	rect := rect{10, 5}
	cir := Circle{10}
	measure(rect)
	measure(cir)

	myPrint(1, 2.5, "Hello", "world")
	printType(1)
	printType(2.4)
	printType("Hello")

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type : int")
	case float64:
		fmt.Println("Type: float64")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: unknown")

	}
}

func myPrint(i ...interface{}) {

	for _, val := range i {
		fmt.Println(val)
	}
}

