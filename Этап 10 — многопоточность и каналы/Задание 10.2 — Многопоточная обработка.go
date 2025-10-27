package main

import "fmt"

func main() {
	numbers := make([]int, 20)
	for i := 0; i < 20; i++ {
		numbers[i] = i + 1
	}

	jobs := make(chan int, 20)
	// Создаем канал результатов с размером буфера = количеству чисел
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
	//close(jobs)

	// Создаем срез для результатов в правильном порядке
	sortedResults := make([]struct {
		num    int
		square int
	}, 20)

	// Собираем результаты и размещаем по правильным индексам
	for i := 0; i < 20; i++ {
		result := <-results
		sortedResults[result.num-1] = result // num-1 = индекс в срезе
	}

	// Выводим отсортированные результаты
	for _, result := range sortedResults {
		fmt.Printf("%d^2 = %d\n", result.num, result.square)
	}
}

func worker(jobs <-chan int, results chan<- struct {
	num    int
	square int
}) {
	for num := range jobs {
		results <- struct {
			num    int
			square int
		}{num,
			num * num}
	}
}
