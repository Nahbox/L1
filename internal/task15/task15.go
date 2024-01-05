package task15

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// Проверяем, что длина строки v больше или равна 100
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		// Обработка случая, когда длина строки меньше 100
		justString = v
	}

	fmt.Println(len(v), len(justString), justString)
}

func createHugeString(size int) string {
	// Создаем строку, состоящую из символа 'A', повторенного size раз
	hugeString := strings.Repeat("A", size)
	return hugeString
}

func Task15() {
	someFunc()
}
