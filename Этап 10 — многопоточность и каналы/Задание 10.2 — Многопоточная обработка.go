package main

import "fmt"

func main() {
	// Создаем массив чисел 1..20
	numbers := make([]int, 20)
	for i := 0; i < 20; i++ {
		numbers[i] = i + 1
	}

	// Создаем каналы
	jobs := make(chan int, 20)
	results := make(chan struct {
		num    int
		square int
	}, 20)

	// Запускаем 4 горутины-воркера
	for i := 0; i < 4; i++ {
		go worker(jobs, results)
	}

	// Отправляем все числа в канал jobs
	for _, num := range numbers {
		jobs <- num
	}

	// Собираем и выводим результаты
	for i := 0; i < 20; i++ {
		result := <-results
		fmt.Printf("%d^2 = %d\n", result.num, result.square)
	}
}

// Функция-воркер
func worker(jobs <-chan int, results chan<- struct {
	num    int
	square int
}) {
	for num := range jobs {
		results <- struct {
			num    int
			square int
		}{num, num * num}
	}
}
