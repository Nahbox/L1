package task26

import "fmt"

// Решение с использованием map для отслеживания использованных символов
func Task26_1() {
	fmt.Println("abcd", areAllCharactersUnique("abcd"))     // true
	fmt.Println("QwErTy", areAllCharactersUnique("QwErTy")) // true
	fmt.Println("abcDe", areAllCharactersUnique("abcDe"))   // true
	fmt.Println("Hello", areAllCharactersUnique("Hello"))   // false
	fmt.Println("Unique", areAllCharactersUnique("Unique")) // false
	fmt.Println("GoLang", areAllCharactersUnique("GoLang")) // false
}

func areAllCharactersUnique(input string) bool {
	charSet := make(map[rune]bool)

	for _, char := range input {
		// Приводим символ к нижнему регистру перед добавлением в множество
		lowercaseChar := toLower(char)

		// Если символ уже присутствует в множестве, строка содержит повторяющиеся символы
		if charSet[lowercaseChar] {
			return false
		}

		charSet[lowercaseChar] = true
	}

	return true
}

func toLower(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		// Приводим символ к нижнему регистру
		return char + ('a' - 'A')
	}
	return char
}
