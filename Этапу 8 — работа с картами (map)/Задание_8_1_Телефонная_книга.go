package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type PhoneBook map[string]string

func main() {
	book := make(PhoneBook)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1 — Добавить контакт")
		fmt.Println("2 — Удалить контакт")
		fmt.Println("3 — Найти контакт")
		fmt.Println("4 — Все контакты")
		fmt.Println("0 — Выход")
		fmt.Print("Выберите пункт: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			name := readNonEmpty(reader, "Введите имя: ")
			phone := readNonEmpty(reader, "Введите телефон: ")

			book[name] = phone
			fmt.Printf("Добавлен контакт: %s -> %s\n", name, phone)
			printAll(book)

		case "2":
			name := readNonEmpty(reader, "Введите имя для удаления: ")
			if _, ok := book[name]; ok {
				delete(book, name)
				fmt.Printf("Удалён контакт: %s\n", name)
			} else {
				fmt.Println("Контакт не найден")
			}
			printAll(book)

		case "3":
			name := readNonEmpty(reader, "Введите имя для поиска: ")
			if phone, ok := book[name]; ok {
				fmt.Printf("Найден контакт: %s -> %s\n", name, phone)
			} else {
				fmt.Println("Контакт не найден")
			}
			printAll(book)

		case "4":
			printAll(book)

		case "0":
			fmt.Println("Выход.")
			return

		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}

func readNonEmpty(r *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		s, _ := r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s != "" {
			return s
		}
		fmt.Println("Поле не может быть пустым.")
	}
}

func printAll(book PhoneBook) {
	if len(book) == 0 {
		fmt.Println("Список пуст.")
		return
	}
	fmt.Println("Все контакты:")
	names := make([]string, 0, len(book))
	for name := range book {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %s\n", name, book[name])
	}
}
