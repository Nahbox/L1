package task26

import (
	"fmt"
	"sort"
	"strings"
)

// Решение: сортируем входную строку и попарно проверяем все символы на идентичность
func Task26_2() {
	fmt.Println("abcd", areAllCharactersUniqueSorted("abcd"))     // true
	fmt.Println("QwErTy", areAllCharactersUniqueSorted("QwErTy")) // true
	fmt.Println("abcDe", areAllCharactersUniqueSorted("abcDe"))   // true
	fmt.Println("Hello", areAllCharactersUniqueSorted("Hello"))   // false
	fmt.Println("Unique", areAllCharactersUniqueSorted("Unique")) // false
	fmt.Println("GoLang", areAllCharactersUniqueSorted("GoLang")) // false
}

func areAllCharactersUniqueSorted(input string) bool {
	// Приводим строку к нижнему регистру перед сортировкой
	sortedString := sortString(strings.ToLower(input))

	// Проверяем наличие повторяющихся символов после сортировки
	for i := 0; i < len(sortedString)-1; i++ {
		if sortedString[i] == sortedString[i+1] {
			return false
		}
	}

	return true
}

func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}
