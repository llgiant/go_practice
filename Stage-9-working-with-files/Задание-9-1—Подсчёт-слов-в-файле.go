package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		words = append(words, fields...)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Количество слов: %d\n", len(words))
	for _, word := range words {
		fmt.Println(word)
	}
}
