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


// Задача 1: Сколько времени выполнится?
func task1() {
	fmt.Println("\n=== Задача 1 ===")
	
	worker := func() <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(500 * time.Millisecond)
			close(ch)
		}()
		return ch
	}
	
	start := time.Now()
	<-worker()
	<-worker()
	<-worker()
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~1.5 секунды (500мс * 3)
}

// Задача 2: Сколько времени выполнится?
func task2() {
	fmt.Println("\n=== Задача 2 ===")
	
	worker := func() <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(1 * time.Second)
			close(ch)
		}()
		return ch
	}
	
	start := time.Now()
	_, _, _ = worker(), worker(), worker()
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~0 секунд (каналы созданы, но не читаем из них)
}

// Задача 3: Сколько времени выполнится?
func task3() {
	fmt.Println("\n=== Задача 3 ===")
	
	worker := func(delay time.Duration) <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(delay)
			close(ch)
		}()
		return ch
	}
	
	start := time.Now()
	<-worker(1 * time.Second)
	<-worker(2 * time.Second)
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~3 секунды (1с + 2с последовательно)
}

// Задача 4: Сколько времени выполнится?
func task4() {
	fmt.Println("\n=== Задача 4 ===")
	
	worker := func() <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(1 * time.Second)
			ch <- 42
		}()
		return ch
	}
	
	start := time.Now()
	ch1 := worker()
	ch2 := worker()
	<-ch1
	<-ch2
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~1 секунда (горутины запускаются параллельно)
}

// Задача 5: Что выведется и сколько времени выполнится?
func task5() {
	fmt.Println("\n=== Задача 5 ===")
	
	worker := func(id int) <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(time.Duration(id) * time.Second)
			ch <- id
		}()
		return ch
	}
	
	start := time.Now()
	val1 := <-worker(2)
	val2 := <-worker(1)
	fmt.Printf("Значения: %d, %d\n", val1, val2)
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: 2, 1; ~3 секунды (2с + 1с)
}

// Задача 6: Сколько времени выполнится?
func task6() {
	fmt.Println("\n=== Задача 6 ===")
	
	worker := func() <-chan int {
		ch := make(chan int, 1) // буферизованный канал
		go func() {
			time.Sleep(1 * time.Second)
			ch <- 42
		}()
		return ch
	}
	
	start := time.Now()
	_, _ = <-worker(), <-worker()
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~2 секунды (последовательное чтение)
}

// Задача 7: Сколько времени выполнится?
func task7() {
	fmt.Println("\n=== Задача 7 ===")
	
	worker := func() <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(1 * time.Second)
			close(ch)
		}()
		return ch
	}
	
	start := time.Now()
	ch := worker()
	<-ch
	<-ch
	<-ch
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~1 секунда (закрытый канал читается мгновенно)
}

// Задача 8: Сколько времени выполнится?
func task8() {
	fmt.Println("\n=== Задача 8 ===")
	
	worker := func(n int) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; i < n; i++ {
				time.Sleep(500 * time.Millisecond)
				ch <- i
			}
			close(ch)
		}()
		return ch
	}
	
	start := time.Now()
	ch := worker(3)
	<-ch
	<-ch
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~1 секунда (2 чтения по 500мс)
}

// Задача 9: Что произойдет?
func task9() {
	fmt.Println("\n=== Задача 9 ===")
	
	worker := func() chan int {
		ch := make(chan int)
		// Нет горутины!
		time.Sleep(1 * time.Second)
		close(ch)
		return ch
	}
	
	start := time.Now()
	_, _ = worker(), worker()
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~2 секунды (последовательное выполнение в основном потоке)
}

// Задача 10: Сколько времени выполнится?
func task10() {
	fmt.Println("\n=== Задача 10 ===")
	
	worker := func() <-chan int {
		ch := make(chan int)
		go func() {
			time.Sleep(2 * time.Second)
			close(ch)
		}()
		return ch
	}
	
	start := time.Now()
	ch1, ch2, ch3 := worker(), worker(), worker()
	<-ch1
	<-ch2
	<-ch3
	fmt.Printf("Время выполнения: %v\n", time.Since(start))
	// Ответ: ~2 секунды (все горутины запущены сразу, ждем параллельно)
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
	
	fmt.Println("\n=== Ответы ===")
	fmt.Println("Задача 1: ~1.5с (последовательное чтение, 500мс * 3)")
	fmt.Println("Задача 2: ~0с (каналы созданы, но не читаем)")
	fmt.Println("Задача 3: ~3с (последовательное чтение, 1с + 2с)")
	fmt.Println("Задача 4: ~1с (каналы созданы параллельно, потом читаем)")
	fmt.Println("Задача 5: ~3с, выведет 2,1 (последовательно 2с + 1с)")
	fmt.Println("Задача 6: ~2с (буфер не помогает при последовательном чтении)")
	fmt.Println("Задача 7: ~1с (закрытый канал читается мгновенно)")
	fmt.Println("Задача 8: ~1с (2 чтения по 500мс каждое)")
	fmt.Println("Задача 9: ~2с (Sleep в основном потоке, нет горутин)")
	fmt.Println("Задача 10: ~2с (все горутины стартуют параллельно)")
}