package task15

import (
	"fmt"
	"strings"
)

// Функция Task15 вызывает функцию someFunc.
func Task15() {
	someFunc()
}

// Функция someFunc выполняет операции с созданием и обработкой большой строки.
func someFunc() {
	// Переменная justString используется для хранения подстроки из большой строки.
	var justString string

	// Создание строки размером 1 КБ, состоящей из символа 'A'.
	v := createHugeString(1 << 10)

	// Проверка, что длина строки v больше или равна 100.
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		// Обработка случая, когда длина строки меньше 100.
		justString = v
	}

	// Вывод длины и содержимого исходной и обрезанной строк.
	fmt.Println(len(v), len(justString), justString)
}

// Функция createHugeString создает строку, состоящую из символа 'A', повторенного size раз.
func createHugeString(size int) string {
	hugeString := strings.Repeat("A", size)
	return hugeString
}
