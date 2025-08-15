package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

type Box struct {
	Rectangle
	Depth float64
}

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func (b Box) Volume() float64 {
	return b.Area() * b.Depth
}

func main() {
	box := Box{Rectangle{2, 3}, 4}
	fmt.Println("Объём:", box.Volume())
}
