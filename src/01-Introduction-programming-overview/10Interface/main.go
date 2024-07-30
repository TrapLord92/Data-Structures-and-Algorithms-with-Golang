package main

import (
	"fmt"
	"math"
)

/*Interface
Interfaces are defined as a set of methods.
Syntax of interface:
Type <Interface name> interface {
<Method name> <Return type>*/

// interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// struct
type Rect struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

// implementation
func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}
func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 10}
	fmt.Println("Total Area: ", TotalArea(r, c))
	fmt.Println("Total Perimeter: ", TotalPerimeter(r, c))
}

/*Analysis:
· A Shape interface is created which contain two methods Area() and Perimeter().
· Rect and Circle implements Shape interface as they implement Area() and Perimeter() methods.
· TotalArea() and TotalPerimeter() are two functions which expect list of object of type Shape.*/
