package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make([]int, 100000000)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// Замер времени для параллельной версии
	startParallel := time.Now()
	parallelSum := calculateParallel(numbers)
	timeParallel := time.Since(startParallel)

	// Замер времени для последовательной версии
	startSequential := time.Now()
	sequentialSum := calculateSequential(numbers)
	timeSequential := time.Since(startSequential)

	fmt.Printf("Параллельная версия:\n")
	fmt.Printf("Сумма = %d, Время = %v\n", parallelSum, timeParallel)

	fmt.Printf("\nПоследовательная версия:\n")
	fmt.Printf("Сумма = %d, Время = %v\n", sequentialSum, timeSequential)

	fmt.Printf("\nПараллельная быстрее в %.2f раз\n",
		float64(timeSequential)/float64(timeParallel))
}

// Параллельная версия с горутинами
func calculateParallel(numbers []int) int {
	firstHalf, secondHalf := devideSlice(numbers)
	sumChan := make(chan int, 2)

	go func() {
		sumChan <- numbersSum(firstHalf)
	}()

	go func() {
		sumChan <- numbersSum(secondHalf)
	}()

	firstHalfSum := <-sumChan
	secondHalfSum := <-sumChan

	return firstHalfSum + secondHalfSum
}

// Последовательная версия
func calculateSequential(numbers []int) int {
	firstHalf, secondHalf := devideSlice(numbers)
	firstHalfSum := numbersSum(firstHalf)
	secondHalfSum := numbersSum(secondHalf)
	return firstHalfSum + secondHalfSum
}

func numbersSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		// Имитация сложных вычислений
		for j := 0; j < 100; j++ {
			sum += num * j / (j + 1)
		}
	}
	return sum
}

func devideSlice(slice []int) ([]int, []int) {
	mid := len(slice) / 2
	return slice[:mid], slice[mid:]
}
