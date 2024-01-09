package task19

import "fmt"

// Функция Task19 демонстрирует обращение строки с использованием руны.
func Task19() {
	// Исходные строки
	str1 := "qwerty"
	str2 := "йцукен"
	str3 := "♨★✪❤✰✈"

	// Обращение строк
	reverseStr1 := reverseString(str1)
	reverseStr2 := reverseString(str2)
	reverseStr3 := reverseString(str3)

	// Вывод результатов
	fmt.Println(str1, "->", reverseStr1)
	fmt.Println(str2, "->", reverseStr2)
	fmt.Println(str3, "->", reverseStr3)
}

// Функция reverseString обращает строку, используя руны.
func reverseString(input string) string {
	// Преобразование строки в срез рун
	runes := []rune(input)

	// Обращение среза рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразование среза рун обратно в строку
	reversed := string(runes)

	return reversed
}
