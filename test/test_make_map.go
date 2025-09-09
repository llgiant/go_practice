package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "dog", "cat", "bird", "dog", "fish"}

	fmt.Println("Исходный список:", words)
	fmt.Println("Создаем пустую map...")

	uniqueMap := make(map[string]struct{})

	fmt.Println("\nПроцесс обработки:")
	fmt.Println("Слово\t\tДействие\t\tТекущие уникальные слова")
	fmt.Println("-----\t\t--------\t\t----------------------")

	for _, word := range words {
		action := "ДОБАВЛЕНО"
		if _, exists := uniqueMap[word]; exists {
			action = "ПРОПУЩЕНО (дубликат)"
		}

		uniqueMap[word] = struct{}{}
		fmt.Printf("%-10s\t%-20s\t%v\n", word, action, getKeysSorted(uniqueMap))
	}

	fmt.Printf("\n✅ Результат: %v\n", getKeysSorted(uniqueMap))
}

func getKeysSorted(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// Сортируем для наглядности
	sort.Strings(keys)
	return keys
}
