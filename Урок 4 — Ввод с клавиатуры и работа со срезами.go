package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateSum(grades []int) int {
	var sum int
	for _, grade := range grades {
		sum += grade
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers []int

	for {
		fmt.Print("Введите число (для выхода введите отрицательное): ")

		// Читаем всю строку (чтобы избежать проблем с буфером)
		inputStr, _ := reader.ReadString('\n')
		inputStr = strings.TrimSpace(inputStr) // Удаляем пробелы и перевод строки

		// Пытаемся преобразовать ввод в число
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Ошибка: введено не число! Попробуйте снова.")
			continue
		}

		// Проверяем условие выхода
		if input < 0 {
			fmt.Println("Выход из цикла.")
			break
		}

		// Добавляем число в срез
		numbers = append(numbers, input)
		fmt.Printf("Добавлено: %d\n", input)
	}

	fmt.Println("Вы ввели: ", numbers)
	fmt.Println("Сумма: ", calculateSum(numbers))
}
