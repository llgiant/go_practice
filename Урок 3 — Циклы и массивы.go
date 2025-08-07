package main

import (
	"fmt"
	"strings"
)

const gradesCount = 5

func calculateAverage(grades [gradesCount]float32) float32 {
	var sum float32
	for _, grade := range grades {
		sum += grade
	}
	return sum / float32(gradesCount)
}

func main() {
	// Исходный массив оценок
	grades := [gradesCount]float32{4.0, 3.7, 5.0, 4.5, 4.2}

	// Форматируем вывод оценок
	var gradesOutput strings.Builder
	gradesOutput.WriteString("Оценки: [")
	for i, grade := range grades {
		if i > 0 {
			gradesOutput.WriteString(" ")
		}
		gradesOutput.WriteString(fmt.Sprintf("%.1f", grade))
	}
	gradesOutput.WriteString("]")

	// Вычисляем и выводим результаты
	average := calculateAverage(grades)
	fmt.Println(gradesOutput.String())
	fmt.Printf("Средний балл: %.2f\n", average)
}
