package main

import (
	"fmt"
)

func main() {
	name := "John"
	age := 30
	height := 1.75
	hasDriverLicense := true

	fmt.Println("Имя: " + name)
	fmt.Println("Возраст:", age)
	fmt.Printf("Рост: %.2f м\n", height)
	if hasDriverLicense {
		fmt.Println("Водительские права: есть")
	} else {
		fmt.Println("Водительские права: нет")
	}
}
