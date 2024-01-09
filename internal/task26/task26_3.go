package task26

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

// Использование структуры Set сторонней библиотеки
func Task26_3() {
	fmt.Println("abcd", areAllCharactersUniqueSet("abcd"))     // true
	fmt.Println("QwErTy", areAllCharactersUniqueSet("QwErTy")) // true
	fmt.Println("abcDe", areAllCharactersUniqueSet("abcDe"))   // true
	fmt.Println("Hello", areAllCharactersUniqueSet("Hello"))   // false
	fmt.Println("Unique", areAllCharactersUniqueSet("Unique")) // false
	fmt.Println("GoLang", areAllCharactersUniqueSet("GoLang")) // false
}

func areAllCharactersUniqueSet(input string) bool {
	charSet := hashset.New()

	for _, char := range strings.ToLower(input) {
		if charSet.Contains(char) {
			return false
		}
		charSet.Add(char)
	}

	return true
}
