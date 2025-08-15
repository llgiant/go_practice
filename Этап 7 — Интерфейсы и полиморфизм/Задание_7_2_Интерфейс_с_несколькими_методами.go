package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func main() {
	shapes := []Shape{
		Circle{10},
		Rectangle{10, 10},
	}

	for _, s := range shapes {
		switch v := s.(type) {
		case Circle:
			fmt.Printf("Это круг радиусом %.2f\n", v.Radius)
		case Rectangle:
			fmt.Printf("Это прямоугольник шириной %.2f и высотой %.2f\n", v.Width, v.Height)
		default:
			fmt.Println("Неизвестная фигура")
		}

		fmt.Printf("Площадь = %.2f, Периметр = %.2f\n", s.Area(), s.Perimeter())
	}
}
