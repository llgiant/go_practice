package main

import "fmt"

func makeCounter() func() int {
	counter := 0

	return func() int {
		counter++
		return counter
	}
}

func main() {
	fn := makeCounter()
	fmt.Println(fn()) // 1
	fmt.Println(fn()) // 2
	fmt.Println(fn()) // 3

	fn2 := makeCounter()
	fmt.Println(fn2()) // 1
	fmt.Println(fn2()) // 2
}
