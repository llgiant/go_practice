package main

import (
	"fmt"
)

func main() {
	var weight float32
	var height float32

	fmt.Print("Type a weight: ")
	fmt.Scan(&weight)
	fmt.Print("Type a height: ")
	fmt.Scan(&height)

	imt := weight / (height * height)
	fmt.Printf("Вес: %.1f кг\n", weight)
	fmt.Printf("Рост: %.2f м\n", height)
	fmt.Printf("ИМТ:: %.2f ", imt)

	if imt < 18.5 {
		fmt.Println("Категория: \"Недовес\"")
	} else if imt >= 18.5 && imt <= 24.9 {
		fmt.Println("Категория: \"Нормальный вес\"")
	} else if imt >= 25 && imt <= 29.9 {
		fmt.Println("Категория: \"Избыточный вес\"")
	} else if imt >= 30 {
		fmt.Println("Категория: \"Ожирение\"")
	}
}
