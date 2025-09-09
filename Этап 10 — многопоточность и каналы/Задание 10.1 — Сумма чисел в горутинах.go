package main

import "fmt"

func main() {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i + 1
	}

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

	totalSum := firstHalfSum + secondHalfSum

	fmt.Printf("Сумма 1..50 = %d\n", firstHalfSum)
	fmt.Printf("Сумма 51..100 = %d\n", secondHalfSum)
	fmt.Printf("Общая сумма = %d\n", totalSum)
}

func numbersSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func devideSlice(slice []int) ([]int, []int) {
	mid := len(slice) / 2
	return slice[:mid], slice[mid:]
}
