package main

import "fmt"

func swap(a, b *int) {

	c := *a
	*a = *b
	*b = c
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
