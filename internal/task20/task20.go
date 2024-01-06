package task20

import (
	"fmt"
	"strings"
)

func Task20() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}

// reverseWords переворачивает слова в строке
func reverseWords(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова
	reversedWords := make([]string, len(words))

	// Переворачиваем каждое слово
	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	// Собираем перевернутую строку из слов
	result := strings.Join(reversedWords, " ")
	return result
}
