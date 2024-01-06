package task19

import "fmt"

func Task19() {
	str1 := "qwerty"
	str2 := "йцукен"
	str3 := "♨★✪❤✰✈"

	reverseStr1 := reverseString(str1)
	reverseStr2 := reverseString(str2)
	reverseStr3 := reverseString(str3)

	fmt.Println(str1, "->", reverseStr1)
	fmt.Println(str2, "->", reverseStr2)
	fmt.Println(str3, "->", reverseStr3)
}

func reverseString(input string) string {
	// Преобразуем строку в срез символов
	runes := []rune(input)

	// Переворачиваем срез символов
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем срез символов обратно в строку
	reversed := string(runes)

	return reversed
}
