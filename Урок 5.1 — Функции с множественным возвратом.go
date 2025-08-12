package main

import "fmt"

func findMinMax(values []int) (int, int) {
	if len(values) == 0 {
		fmt.Println("Срез пустой!")
		return 0, 0
	}

	min, max := values[0], values[0]

	for _, value := range values[1:] {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	//values := []int{3, 8, 2, 5, 1, 9}
	values := []int{} // ← попробуй с пустым срезом

	min, max := findMinMax(values)
	fmt.Printf("Минимум: %d\n", min)
	fmt.Printf("Максимум: %d\n", max)
}
