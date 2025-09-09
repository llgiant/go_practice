package main

import (
	"bufio"
	"os"
	"sort"
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

	uniqueMap := make(map[string]struct{})

	for _, word := range words {
		uniqueMap[word] = struct{}{}
	}

	uniqueWords := make([]string, 0, len(uniqueMap))

	for word := range uniqueMap {
		uniqueWords = append(uniqueWords, word)
	}

	sort.Strings(uniqueWords)

	outputFile, err := os.Create("words.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, word := range uniqueWords {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

	if err := writer.Flush(); err != nil {
		panic(err)
	}
}
