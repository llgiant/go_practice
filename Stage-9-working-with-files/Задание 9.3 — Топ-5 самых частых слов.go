package main

import (
	"bufio"
	"fmt"
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
		for i := range fields {
			fields[i] = strings.ToLower(fields[i])
		}
		words = append(words, fields...)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	wordCounter := make(map[string]int)
	for _, word := range words {
		wordCounter[word]++
	}

	type WordCount struct {
		Word  string
		Count int
	}

	wordCounts := make([]WordCount, 0, len(wordCounter))

	for word, count := range wordCounter {
		wordCounts = append(wordCounts, WordCount{word, count})
	}

	sort.Slice(wordCounts,
		func(i, j int) bool {
			if wordCounts[i].Count == wordCounts[j].Count {
				return wordCounts[i].Word < wordCounts[j].Word
			}
			return wordCounts[i].Count > wordCounts[j].Count
		})
	limit := 5
	if len(wordCounts) < limit {
		limit = len(wordCounts)
	}

	for _, wc := range wordCounts[:limit] {
		fmt.Printf("%-15s: %d\n", wc.Word, wc.Count)
	}
}
