package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func main() {
	shapes := []Shape{
		Rectangle{3, 5},
		Circle{5},
	}

	for _, s := range shapes {
		// Печатаем имя типа и площадь с двумя знаками после запятой
		fmt.Printf("%T: %.2f\n", s, s.Area())
	}
}
