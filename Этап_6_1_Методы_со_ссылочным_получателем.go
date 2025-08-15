package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
func main() {

	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Исходная площадь:", rect.Area()) // Площадь: 15

	rect.Scale(2)
	fmt.Println("Новая площадь:", rect.Area()) // Площадь: 60

}
