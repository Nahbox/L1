package task20

import (
	"fmt"
	"strings"
)

// Функция Task20 демонстрирует переворачивание слов в строке.
func Task20() {
	// Исходная строка
	input := "snow dog sun"
	// Перевернутая строка с перевернутыми словами
	reversed := reverseWords(input)
	fmt.Println(reversed)
}

// Функция reverseWords переворачивает слова в строке.
func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)
	reversedWords := make([]string, len(words))

	// Переворачиваем каждое слово
	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	// Собираем перевернутую строку из слов
	result := strings.Join(reversedWords, " ")
	return result
}
