package main

import "fmt"

func main() {
	i := 24.4
	var pf *float64

	pf = &i
	fmt.Println("Value:", pf)
	fmt.Println("Value:", *pf)
}
