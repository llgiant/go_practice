package main

import (
	"fmt"
	"time"
)

//
//func worker() <-chan int {
//	result := make(chan int)
//	go func() {
//		time.Sleep(1 * time.Second)
//		close(result)
//	}()
//
//	return result
//}

//func main() {
//	timeStart := time.Now()
//	_, _, _ = <-worker(), <-worker(), <-worker()
//	fmt.Println(time.Since(timeStart))
//	newtimeStart := time.Now()
//	_, _ = worker(), worker()
//	fmt.Println(time.Since(newtimeStart))
//	task8()
//}

func main() {
	fmt.Println("\n=== Задача 8 ===")

	worker := func(n int) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; i < n; i++ {
				time.Sleep(500 * time.Millisecond)
				ch <- n
			}
			close(ch)
		}()
		return ch
	}

	start := time.Now()
	val1 := <-worker(3)
	val2 := <-worker(300)
	fmt.Printf("Значения: %d, %d\n", val1, val2)

	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~1 секунда (2 чтения по 500мс)
}
